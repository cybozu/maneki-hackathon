package main

import (
	"bytes"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"os"
)

// main.go indexName
func main() {
	es, _ := elasticsearch.NewDefaultClient()
	indexName := os.Args[1]
	res, err := es.Indices.Delete([]string{indexName})
	if err != nil {
		fmt.Printf("error occured %s", err)
		return
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	fmt.Printf("Delete OK: %s\n", buf.String())
	if res, err := es.Indices.Create(indexName); err == nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		fmt.Printf("Create OK: %s\n", buf.String())
	}
	reader, _ := os.Open("./document.json")
	defer reader.Close()
	if res, err := es.Index(indexName, reader); err == nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		fmt.Printf("PUT document OK: %s\n", buf.String())
	}
	if res, err := es.Cat.Indices(); err == nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		fmt.Printf("Cat OK: %s\n", buf.String())
	}
	if res, err := es.Indices.Get([]string{indexName}); err == nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		fmt.Printf("Get document OK: %s\n", buf.String())
	}
}
