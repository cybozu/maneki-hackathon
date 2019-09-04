package main

import (
	"fmt"
	"bytes"
	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	es, _ := elasticsearch.NewDefaultClient()
	if res, err := es.Cat.Indices(); err == nil {	
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		fmt.Printf("OK: %s\n", buf.String())
	}
}
