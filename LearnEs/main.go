package main

import (
	"context"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/densevectorsimilarity"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/indices/getmapping"
)

func duibi() {
	client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		CloudID: "My_Elasticsearch_project:dXMtZWFzdC0xLmF3cy5lbGFzdGljLmNsb3VkJGRkY2Y0ZjY0NzRkYjRkOGU5MWEwZWI2YTEzODY4ZDAyLmVzJGRkY2Y0ZjY0NzRkYjRkOGU5MWEwZWI2YTEzODY4ZDAyLmti",
		APIKey:  "bG5WZExKY0JZOXgtVk9YdnlrbEo6R0RzTkl5T0E0YVdPYjJlaWF3UEtsUQ==",
	})
	if err != nil {
		panic(err)
	}

	req := getmapping.New(client.Transport)
	req.Index("my_es9")
	res, err := req.Do(context.Background())
	if err != nil {
		log.Fatalf("Error getting mapping: %s", err)
	}

	dims := 2560
	index := true
	similarity := densevectorsimilarity.Cosine

	mapping := types.TypeMapping{
		Properties: map[string]types.Property{
			"content":       types.NewTextProperty(),
			"extralocation": types.NewTextProperty(),
			"content_dense_vector": &types.DenseVectorProperty{
				Dims:       &dims,
				Index:      &index,
				Similarity: &similarity,
			},
		},
	}

	remoteMapping := res["my_es9"].Mappings.Properties

	//对比remoteMapping和mapping，如果remoteMapping和mappingkey不匹配，则报错
	for k, _ := range remoteMapping {
		if _, ok := mapping.Properties[k]; !ok {
			log.Fatalf("remoteMapping and mapping key not match: %s", k)
		}
	}

	// 打印返回的 mapping
	fmt.Printf("Mapping for index 'your-index-name':\n%v\n", res["my_es9"].Mappings.Properties)
	// Create Index
	// res, err := client.Indices.Create("first_try").Do(context.Background())
	// if err != nil {
	// 	log.Panicf("create index err:%v", err)
	// }
	// fmt.Println("Create Index Succefully", res)

	//index a document
	// document := struct {
	// 	Name string `json:"name"`
	// }{
	// 	Name: "woodQQQ",
	// }
	// res, err := client.Index("first_try").Id("1").Request(document).Do(context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Index Document Succefully", res)

	//get a document
	// res, err := client.Get("first_try", "1").Do(context.Background())
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println("Get Document Succefully", string(res.Source_))

	//search document
	// res, err := client.Search().
	// 	Index("index_name").
	// 	Request(&search.Request{
	// 		Query: &types.Query{
	// 			Match: map[string]types.MatchQuery{
	// 				"name": {Query: "Foo"},
	// 			},
	// 		},
	// 	}).Do(context.Background())
	// if err != nil {
	// 	log.Panic(err)
	// }
	// for _, hit := range res.Hits.Hits {
	// 	fmt.Println(string(hit.Source_))
	// }

	//update document
	// res, err := client.Update("first_try", "1").Request(&update.Request{
	// 	Doc: json.RawMessage(`{ "name" : "woodQ" }`),
	// }).Do(context.Background())
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println(res.Result)

	//Delete index
	//typedClient.Indices.Delete("my_index").Do(context.TODO())

	//Delete document
	//typedClient.Delete("my_index", "id").Do(context.TODO())
}
