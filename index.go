package main

import (
	"bytes"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
)

func createIndex(indexName string) string {
	es, _ := elasticsearch.NewDefaultClient()
	res, err := es.Indices.Create(indexName)
	if err != nil {
		fmt.Printf("error creating index: %s", err)
		return "error"
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	fmt.Printf("Create OK: %s\n", buf.String())
	return "success"
}
