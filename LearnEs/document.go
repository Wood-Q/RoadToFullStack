package main

import (
	"context"
	"fmt"
	"log"

	"learnMQ/LearnEs/es8"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino/schema"
	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	ctx := context.Background()
	mapping := `{
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
					"dims": 2580,
					"index": true,
					"similarity": "cosine"
				}
			}
		}
	}`
	// 创建 embedding 组件
	embedder, err := ark.NewEmbedder(context.Background(), &ark.EmbeddingConfig{
		APIKey: "56a6b406-8b6b-4bb5-b169-92117a5caa72",
		Model:  "doubao-embedding-text-240715",
	})
	if err != nil {
		log.Panicf("create embedder failed, err=%v", err)
	}
	// 1. 创建ES客户端
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		CloudID: "My_Elasticsearch_project:dXMtZWFzdC0xLmF3cy5lbGFzdGljLmNsb3VkJGRkY2Y0ZjY0NzRkYjRkOGU5MWEwZWI2YTEzODY4ZDAyLmVzJGRkY2Y0ZjY0NzRkYjRkOGU5MWEwZWI2YTEzODY4ZDAyLmti",
		APIKey:  "bG5WZExKY0JZOXgtVk9YdnlrbEo6R0RzTkl5T0E0YVdPYjJlaWF3UEtsUQ==",
	})
	if err != nil {
		log.Fatal(err)
	}

	// 2. 确保索引存在
	indexName := "search-e009"

	// 5. 创建索引器
	indexer, err := es8.NewIndexer(ctx, &es8.IndexerConfig{
		Client:    client,
		Index:     indexName,
		BatchSize: 5,
		Embedding: embedder,
		DocumentToFields: func(ctx context.Context, doc *schema.Document) (field2Value map[string]es8.FieldValue, err error) {
			return map[string]es8.FieldValue{
				"id": {
					Value: doc.ID,
				},
				"title": {
					Value: doc.MetaData["title"],
				},
				"content": {
					Value:    doc.Content,
					EmbedKey: "content_dense_vector", // vectorize doc content and save vector to field "content_vector"
				},
			}, nil
		},
		CustomMapping: mapping,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 6. 准备要存储的文档
	docs := []*schema.Document{
		{
			ID:      "doc_4",
			Content: "这是关",
		},
		{
			ID:      "doc_5",
			Content: "学习的文档内容",
			MetaData: map[string]any{
				"title": "机器基础",
			},
		},
		{
			ID:      "doc_6",
			Content: "这是关于深档内容",
			MetaData: map[string]any{
				"title": "深度原理",
			},
		},
	}

	// 7. 实际存储文档到ES
	fmt.Println("开始存储文档...")
	ids, err := indexer.Store(ctx, docs)
	if err != nil {
		log.Fatal("存储文档失败:", err)
	}

	fmt.Printf("成功存储 %d 个文档，IDs: %v\n", len(ids), ids)

	// 8. 验证存储结果
	if err := verifyStorage(client, indexName, ids); err != nil {
		log.Printf("验证存储结果失败: %v", err)
	} else {
		fmt.Println("所有文档已成功存储到Elasticsearch!")
	}
}

// 验证文档是否成功存储
func verifyStorage(client *elasticsearch.Client, indexName string, docIDs []string) error {
	for _, id := range docIDs {
		res, err := client.Get(indexName, id)
		if err != nil {
			return fmt.Errorf("获取文档 %s 失败: %w", id, err)
		}
		if res.StatusCode == 404 {
			return fmt.Errorf("文档 %s 未找到", id)
		}
		fmt.Printf("✓ 文档 %s 存储成功\n", id)
	}
	return nil
}
