package main

import (
	"context"
	"fmt"
	"log"

	"learnMQ/LearnEs/es8"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino/schema"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/densevectorsimilarity"
)

func main() {
	ctx := context.Background()
	dims := 2560
	similarity := densevectorsimilarity.Cosine
	index := true
	//创建本地mapping
	mapping := &types.TypeMapping{
		Properties: map[string]types.Property{
			"id":             types.NewLongNumberProperty(),
			"title":          types.NewTextProperty(),
			"content":        types.NewTextProperty(),
			"extra_location": types.NewTextProperty(),
			"content_dense_vector": &types.DenseVectorProperty{
				Dims:       &dims,
				Index:      &index,
				Similarity: &similarity,
			},
		},
	}
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
	indexName := "my_es9"

	// 5. 创建索引器
	indexer, err := es8.NewIndexer(ctx, &es8.IndexerConfig{
		Client:       client,
		Index:        indexName,
		BatchSize:    5,
		LocalMapping: mapping,
		Embedding:    embedder,
		DocumentToFields: func(ctx context.Context, doc *schema.Document) (field2Value map[string]es8.FieldValue, err error) {
			return map[string]es8.FieldValue{
				"id": {
					Value: doc.ID,
				},
				"content": {
					Value:    doc.Content,
					EmbedKey: "content_dense_vector",
				},
			}, nil
		},
		ValidationMode:    es8.ValidationModeWarn,
		EnableSchemaCheck: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 6. 准备要存储的文档
	docs := []*schema.Document{
		{
			ID:      "doc_1001",
			Content: "doc_1001",
		},
	}

	// 7. 实际存储文档到ES
	fmt.Println("开始存储文档...")
	ids, err := indexer.Store(ctx, docs)
	if err != nil {
		log.Fatal("存储文档失败:", err)
	}

	fmt.Printf("成功存储 %d 个文档，IDs: %v\n", len(ids), ids)
}
