package main

import (
	"context"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v9"
)

func ma() {
	client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		Username:  "elastic",
		Password:  "LUB7jymz",
	})
	if err != nil {
		panic(err)
	}
	// Create Index
	res, err := client.Indices.Create("first_try").Do(context.Background())
	if err != nil {
		log.Panicf("create index err:%v", err)
	}
	fmt.Println("Create Index Succefully", res)

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
