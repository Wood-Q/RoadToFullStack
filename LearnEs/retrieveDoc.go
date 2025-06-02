package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino-ext/components/retriever/es8"
	"github.com/cloudwego/eino-ext/components/retriever/es8/search_mode"
	"github.com/cloudwego/eino/schema"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func main() {
	ctx := context.Background()
	// 1. 创建ES客户端
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		CloudID: "My_Elasticsearch_project:dXMtZWFzdC0xLmF3cy5lbGFzdGljLmNsb3VkJGRkY2Y0ZjY0NzRkYjRkOGU5MWEwZWI2YTEzODY4ZDAyLmVzJGRkY2Y0ZjY0NzRkYjRkOGU5MWEwZWI2YTEzODY4ZDAyLmti",
		APIKey:  "bG5WZExKY0JZOXgtVk9YdnlrbEo6R0RzTkl5T0E0YVdPYjJlaWF3UEtsUQ==",
	})
	if err != nil {
		log.Fatal(err)
	}

	// 创建 embedding 组件
	embedder, err := ark.NewEmbedder(context.Background(), &ark.EmbeddingConfig{
		APIKey: "56a6b406-8b6b-4bb5-b169-92117a5caa72",
		Model:  "doubao-embedding-text-240715",
	})
	if err != nil {
		log.Panicf("create embedder failed, err=%v", err)
	}

	retriever, err := es8.NewRetriever(ctx, &es8.RetrieverConfig{
		Client: client,
		Index:  "my_rag",
		SearchMode: search_mode.SearchModeApproximate(&search_mode.ApproximateConfig{
			QueryFieldName:  "content",
			VectorFieldName: "content_dense_vector",
			Hybrid:          true,
			// RRF only available with specific licenses
			// see: https://www.elastic.co/subscriptions
			RRF:             false,
			RRFRankConstant: nil,
			RRFWindowSize:   nil,
		}),
		ResultParser: func(ctx context.Context, hit types.Hit) (doc *schema.Document, err error) {
			doc = &schema.Document{
				ID:       *hit.Id_,
				Content:  "",
				MetaData: map[string]any{},
			}

			var src map[string]any
			if err = json.Unmarshal(hit.Source_, &src); err != nil {
				return nil, err
			}

			for field, val := range src {
				switch field {
				case "content":
					doc.Content = val.(string)
				case "content_dense_vector":
					var v []float64
					for _, item := range val.([]interface{}) {
						v = append(v, item.(float64))
					}
					doc.WithDenseVector(v)
				}
			}

			if hit.Score_ != nil {
				doc.WithScore(float64(*hit.Score_))
			}

			return doc, nil
		},
		Embedding: embedder,
	})
	if err != nil {
		log.Fatal(err)
	}

	docs, err := retriever.Retrieve(ctx, "明月班")
	if err != nil {
		log.Fatal(err)
	}

	for _, doc := range docs {
		fmt.Println(doc.Content)
	}
}
