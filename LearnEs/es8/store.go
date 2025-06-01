/*
 * Copyright 2024 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package es8

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components"
	"github.com/cloudwego/eino/components/embedding"
	"github.com/cloudwego/eino/components/indexer"
	"github.com/cloudwego/eino/schema"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

const defaultBatchSize = 5
const typ = "es8_indexer"

type IndexerConfig struct {
	Client *elasticsearch.Client `json:"client"`

	Index string `json:"index"`
	// BatchSize controls max texts size for embedding.
	// Default is 5.
	BatchSize int `json:"batch_size"`
	// FieldMapping supports customize es fields from eino document.
	// Each key - FieldValue.Value from field2Value will be saved, and
	// vector of FieldValue.Value will be saved if FieldValue.EmbedKey is not empty.
	DocumentToFields func(ctx context.Context, doc *schema.Document) (field2Value map[string]FieldValue, err error)
	// Embedding vectorization method, must provide in two cases
	// 1. VectorFields contains fields except doc Content
	// 2. VectorFields contains doc Content and vector not provided in doc extra (see Document.Vector method)
	Embedding embedding.Embedder
	// CustomMapping allows users to provide custom index mapping
	// If not provided, default mapping will be used
	CustomMapping string `json:"custom_mapping"`
}

type FieldValue struct {
	// Value original Value
	Value any
	// EmbedKey if set, Value will be vectorized and saved to es.
	// If Stringify method is provided, Embedding input text will be Stringify(Value).
	// If Stringify method not set, retriever will try to assert Value as string.
	EmbedKey string
	// Stringify converts Value to string
	Stringify func(val any) (string, error)
}

type Indexer struct {
	client *elasticsearch.Client
	config *IndexerConfig
}

// getDefaultMapping returns the default index mapping
func getDefaultMapping() string {
	return `{
		"mappings": {
			"properties": {
				"content": {
					"type": "text"
				},
				"extra_location": {
					"type": "text"
				},
				"content_dense_vector": {
					"type": "dense_vector",
					"dims": 2560,
					"index": true,
					"similarity": "cosine"
				}
			}
		}
	}`
}

// createIndexIfNotExists 检查索引是否存在，不存在则创建默认索引
func (i *Indexer) createIndexIfNotExists(ctx context.Context) error {
	indexName := i.config.Index

	// 1. 检查索引是否存在
	res, err := i.client.Indices.Exists([]string{indexName})
	if err != nil {
		return fmt.Errorf("failed to check if index exists: %w", err)
	}

	if res.StatusCode == 404 {
		// 2. 索引不存在，创建默认索引
		fmt.Printf("索引 %s 不存在，正在创建默认索引...\n", indexName)

		// 使用默认mapping创建索引
		defaultMapping := getDefaultMapping()

		createRes, err := i.client.Indices.Create(
			indexName,
			i.client.Indices.Create.WithBody(strings.NewReader(defaultMapping)),
			i.client.Indices.Create.WithContext(ctx),
		)
		if err != nil {
			return fmt.Errorf("failed to create index %s: %w", indexName, err)
		}
		defer createRes.Body.Close()

		if createRes.IsError() {
			var createError map[string]interface{}
			if err := json.NewDecoder(createRes.Body).Decode(&createError); err == nil {
				return fmt.Errorf("failed to create index %s: %v", indexName, createError)
			}
			return fmt.Errorf("failed to create index %s: %s", indexName, createRes.String())
		}

		fmt.Printf("✓ 成功创建默认索引 %s\n", indexName)
		return nil
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("unexpected response when checking index existence: %s", res.String())
	}

	// 3. 索引存在，检查用户是否提供了自定义mapping
	if i.config.CustomMapping != "" {
		fmt.Printf("索引 %s 已存在，正在验证mapping一致性...\n", indexName)
		if err := i.validateIndexMapping(ctx); err != nil {
			return fmt.Errorf("索引mapping不一致: %w", err)
		}
		fmt.Printf("✓ 索引 %s mapping验证通过\n", indexName)
	} else {
		fmt.Printf("✓ 索引 %s 已存在，未提供自定义mapping，跳过验证\n", indexName)
	}

	return nil
}

// validateIndexMapping 验证现有索引mapping与用户自定义mapping是否一致
func (i *Indexer) validateIndexMapping(ctx context.Context) error {

	// 获取现有索引的mapping
	currentProperties, err := i.getCurrentIndexMapping(ctx)
	if err != nil {
		return fmt.Errorf("获取当前索引mapping失败: %w", err)
	}

	// 解析用户自定义mapping
	expectedProperties, err := i.parseCustomMapping()
	if err != nil {
		return fmt.Errorf("解析自定义mapping失败: %w", err)
	}

	// 比较mapping结构
	return i.compareMappings(currentProperties, expectedProperties)
}

// getCurrentIndexMapping 获取当前索引的mapping
func (i *Indexer) getCurrentIndexMapping(ctx context.Context) (map[string]interface{}, error) {
	indexName := i.config.Index

	res, err := i.client.Indices.GetMapping(
		i.client.Indices.GetMapping.WithIndex(indexName),
		i.client.Indices.GetMapping.WithContext(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get mapping for index %s: %w", indexName, err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("failed to get mapping for index %s: %s", indexName, res.String())
	}

	var mappingResponse map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&mappingResponse); err != nil {
		return nil, fmt.Errorf("failed to decode mapping response: %w", err)
	}

	// 提取当前mapping的properties
	indexMapping, ok := mappingResponse[indexName].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid mapping response structure for index %s", indexName)
	}

	mappings, ok := indexMapping["mappings"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("no mappings found for index %s", indexName)
	}

	properties, ok := mappings["properties"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("no properties found in mappings for index %s", indexName)
	}

	fmt.Println("当前远程mapping为：", properties)

	return properties, nil
}

// parseCustomMapping 解析用户自定义mapping
func (i *Indexer) parseCustomMapping() (map[string]interface{}, error) {
	var customMappingStruct map[string]interface{}
	if err := json.Unmarshal([]byte(i.config.CustomMapping), &customMappingStruct); err != nil {
		return nil, fmt.Errorf("failed to parse custom mapping: %w", err)
	}

	mappings, ok := customMappingStruct["mappings"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid custom mapping structure: no 'mappings' field")
	}

	properties, ok := mappings["properties"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid custom mapping structure: no 'properties' field in mappings")
	}

	return properties, nil
}

// compareMappings 比较两个mapping的properties是否一致
func (i *Indexer) compareMappings(currentProperties, expectedProperties map[string]interface{}) error {
	// 检查自定义mapping中定义的所有字段是否在当前索引中存在且类型匹配
	for fieldName, expectedField := range expectedProperties {
		expectedFieldMap, ok := expectedField.(map[string]interface{})
		if !ok {
			continue // 跳过无效的字段定义
		}

		// 检查字段是否存在
		currentField, exists := currentProperties[fieldName]
		if !exists {
			return fmt.Errorf("字段 '%s' 在当前索引中不存在，但在自定义mapping中有定义", fieldName)
		}

		currentFieldMap, ok := currentField.(map[string]interface{})
		if !ok {
			return fmt.Errorf("字段 '%s' 的结构无效", fieldName)
		}

		// 比较字段类型
		expectedType, hasExpectedType := expectedFieldMap["type"]
		currentType, hasCurrentType := currentFieldMap["type"]

		if hasExpectedType && hasCurrentType {
			if expectedType != currentType {
				return fmt.Errorf("字段 '%s' 类型不匹配: 期望类型=%v, 当前类型=%v",
					fieldName, expectedType, currentType)
			}
		}

		// 对于dense_vector类型，检查维度是否匹配
		if expectedType == "dense_vector" {
			expectedDims, hasExpectedDims := expectedFieldMap["dims"]
			currentDims, hasCurrentDims := currentFieldMap["dims"]

			if hasExpectedDims && hasCurrentDims {
				if expectedDims != currentDims {
					return fmt.Errorf("字段 '%s' 向量维度不匹配: 期望维度=%v, 当前维度=%v",
						fieldName, expectedDims, currentDims)
				}
			}
		}

		// 可以根据需要添加更多字段属性的检查，如analyzer等
	}

	// 可选：检查当前索引是否有自定义mapping中未定义的字段（警告级别）
	for fieldName := range currentProperties {
		if _, exists := expectedProperties[fieldName]; !exists {
			fmt.Printf("警告: 当前索引中存在字段 '%s'，但在自定义mapping中未定义\n", fieldName)
		}
	}

	return nil
}
func NewIndexer(ctx context.Context, conf *IndexerConfig) (*Indexer, error) {
	if conf.Client == nil {
		return nil, fmt.Errorf("[NewIndexer] es client not provided")
	}

	if conf.DocumentToFields == nil {
		return nil, fmt.Errorf("[NewIndexer] DocumentToFields method not provided")
	}

	if conf.Index == "" {
		return nil, fmt.Errorf("[NewIndexer] index name cannot be empty")
	}

	if conf.BatchSize <= 0 {
		conf.BatchSize = defaultBatchSize
	}

	// Validate batch size is reasonable
	if conf.BatchSize > 1000 {
		return nil, fmt.Errorf("[NewIndexer] batch size %d is too large, maximum recommended is 1000", conf.BatchSize)
	}

	// Test connection to Elasticsearch
	if _, err := conf.Client.Info(); err != nil {
		return nil, fmt.Errorf("[NewIndexer] failed to connect to Elasticsearch: %w", err)
	}

	indexer := &Indexer{
		client: conf.Client,
		config: conf,
	}

	// Create index if not exists or validate mapping
	if err := indexer.createIndexIfNotExists(ctx); err != nil {
		return nil, fmt.Errorf("[NewIndexer] index setup failed: %w", err)
	}

	return indexer, nil
}

func (i *Indexer) Store(ctx context.Context, docs []*schema.Document, opts ...indexer.Option) (ids []string, err error) {
	ctx = callbacks.EnsureRunInfo(ctx, i.GetType(), components.ComponentOfIndexer)
	ctx = callbacks.OnStart(ctx, &indexer.CallbackInput{Docs: docs})
	defer func() {
		if err != nil {
			callbacks.OnError(ctx, err)
		}
	}()

	options := indexer.GetCommonOptions(&indexer.Options{
		Embedding: i.config.Embedding,
	}, opts...)

	if err = i.bulkAdd(ctx, docs, options); err != nil {
		return nil, err
	}

	ids = iter(docs, func(t *schema.Document) string { return t.ID })

	callbacks.OnEnd(ctx, &indexer.CallbackOutput{IDs: ids})

	return ids, nil
}

func (i *Indexer) bulkAdd(ctx context.Context, docs []*schema.Document, options *indexer.Options) error {
	emb := options.Embedding
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         i.config.Index,
		Client:        i.client,
		NumWorkers:    min(4, len(docs)),
		FlushBytes:    int(5e+6),
		FlushInterval: 30 * time.Second,
	})
	if err != nil {
		return err
	}

	var (
		tuples []tuple
		texts  []string
	)

	embAndAdd := func() error {
		var vectors [][]float64

		if len(texts) > 0 {
			if emb == nil {
				return fmt.Errorf("[bulkAdd] embedding method not provided")
			}

			vectors, err = emb.EmbedStrings(i.makeEmbeddingCtx(ctx, emb), texts)
			if err != nil {
				return fmt.Errorf("[bulkAdd] embedding failed, %w", err)
			}

			if len(vectors) != len(texts) {
				return fmt.Errorf("[bulkAdd] invalid vector length, expected=%d, got=%d", len(texts), len(vectors))
			}
		}

		for _, t := range tuples {
			fields := make(map[string]any)

			// Copy original fields
			for k, v := range t.fields {
				fields[k] = v
			}

			// Add vector fields
			for k, idx := range t.key2Idx {
				if idx < len(vectors) {
					fields[k] = vectors[idx]
				}
			}

			b, err := json.Marshal(fields)
			if err != nil {
				return fmt.Errorf("[bulkAdd] marshal bulk item failed, %w", err)
			}

			if err = bi.Add(ctx, esutil.BulkIndexerItem{
				Index:      i.config.Index,
				Action:     "index",
				DocumentID: t.id,
				Body:       bytes.NewReader(b),
				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
					fmt.Printf("✓ 文档 %s 索引成功\n", item.DocumentID)
				},
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						fmt.Printf("✗ 文档 %s 索引失败: %v\n", item.DocumentID, err)
					} else {
						fmt.Printf("✗ 文档 %s 索引失败: %s\n", item.DocumentID, res.Error.Reason)
					}
				},
			}); err != nil {
				return fmt.Errorf("[bulkAdd] failed to add document %s to bulk indexer: %w", t.id, err)
			}
		}

		tuples = tuples[:0]
		texts = texts[:0]

		return nil
	}

	for idx := range docs {
		doc := docs[idx]
		fields, err := i.config.DocumentToFields(ctx, doc)
		if err != nil {
			return fmt.Errorf("[bulkAdd] FieldMapping failed, %w", err)
		}

		rawFields := make(map[string]any)
		embSize := 0
		embedKeys := make(map[string]bool) // Track embed keys to prevent duplicates

		for k, v := range fields {
			rawFields[k] = v.Value
			if v.EmbedKey != "" {
				// Check if EmbedKey conflicts with existing field names
				if _, found := fields[v.EmbedKey]; found {
					return fmt.Errorf("[bulkAdd] embed key '%s' conflicts with existing field for document %s", v.EmbedKey, doc.ID)
				}

				// Check if EmbedKey conflicts with other embed keys
				if embedKeys[v.EmbedKey] {
					return fmt.Errorf("[bulkAdd] duplicate embed key '%s' found for document %s", v.EmbedKey, doc.ID)
				}
				embedKeys[v.EmbedKey] = true

				// Check if EmbedKey would overwrite a raw field
				if _, found := rawFields[v.EmbedKey]; found {
					return fmt.Errorf("[bulkAdd] embed key '%s' would overwrite existing field value for document %s", v.EmbedKey, doc.ID)
				}

				embSize++
			}
		}

		if embSize > i.config.BatchSize {
			return fmt.Errorf("[bulkAdd] needEmbeddingFields length over batch size, batch size=%d, got size=%d",
				i.config.BatchSize, embSize)
		}

		if len(texts)+embSize > i.config.BatchSize {
			if err = embAndAdd(); err != nil {
				return err
			}
		}

		key2Idx := make(map[string]int, embSize)
		for k, v := range fields {
			if v.EmbedKey != "" {
				if _, found := fields[v.EmbedKey]; found {
					return fmt.Errorf("[bulkAdd] duplicate key for origin key, key=%s", k)
				}

				if _, found := key2Idx[v.EmbedKey]; found {
					return fmt.Errorf("[bulkAdd] duplicate key from embed_key, key=%s", v.EmbedKey)
				}

				var text string
				if v.Stringify != nil {
					text, err = v.Stringify(v.Value)
					if err != nil {
						return err
					}
				} else {
					var ok bool
					text, ok = v.Value.(string)
					if !ok {
						return fmt.Errorf("[bulkAdd] assert value as string failed, key=%s, emb_key=%s", k, v.EmbedKey)
					}
				}

				key2Idx[v.EmbedKey] = len(texts)
				texts = append(texts, text)
			}
		}

		tuples = append(tuples, tuple{
			id:      doc.ID,
			fields:  rawFields,
			key2Idx: key2Idx,
		})
	}

	if len(tuples) > 0 {
		if err = embAndAdd(); err != nil {
			return err
		}
	}

	// Close bulk indexer and check statistics
	if err := bi.Close(ctx); err != nil {
		return fmt.Errorf("[bulkAdd] failed to close bulk indexer: %w", err)
	}

	// Check bulk indexer statistics
	stats := bi.Stats()
	fmt.Printf("批量索引完成: 成功=%d, 失败=%d\n", stats.NumIndexed, stats.NumFailed)

	if stats.NumFailed > 0 {
		return fmt.Errorf("[bulkAdd] %d documents failed to index", stats.NumFailed)
	}

	return nil
}

func (i *Indexer) makeEmbeddingCtx(ctx context.Context, emb embedding.Embedder) context.Context {
	runInfo := &callbacks.RunInfo{
		Component: components.ComponentOfEmbedding,
	}

	if embType, ok := components.GetType(emb); ok {
		runInfo.Type = embType
	}

	runInfo.Name = runInfo.Type + string(runInfo.Component)

	return callbacks.ReuseHandlers(ctx, runInfo)
}

func (i *Indexer) GetType() string {
	return typ
}

func (i *Indexer) IsCallbacksEnabled() bool {
	return true
}

type tuple struct {
	id      string
	fields  map[string]any
	key2Idx map[string]int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func iter(docs []*schema.Document, fn func(*schema.Document) string) []string {
	result := make([]string, len(docs))
	for i, doc := range docs {
		result[i] = fn(doc)
	}
	return result
}
