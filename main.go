package main

import (
	"bytes"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	es, _ := elasticsearch.NewDefaultClient()
	indexName := "sample"
	if res, err := es.Indices.Delete([]string{indexName}); err == nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		fmt.Printf("Delete OK: %s\n", buf.String())
	}
	if res, err := es.Indices.Create(indexName); err == nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		fmt.Printf("Create OK: %s\n", buf.String())
	}
	if res, err := es.Cat.Indices(); err == nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		fmt.Printf("Cat OK: %s\n", buf.String())
	}
}
