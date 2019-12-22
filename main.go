package main

import (
	"fmt"
	"github.com/pco2699/hackernews1200/cmd"
)

func main() {
	docs, err := cmd.Fetch()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(docs)
	texts, err := cmd.Extract(docs)
	if err != nil {
		fmt.Println(err.Error())
	}
	tokens, err := cmd.Tokenize(texts)
	if err != nil {
		fmt.Println(err.Error())
	}
	items := cmd.Count(tokens)
	for _, item := range items {
		fmt.Printf("Text: %v Count: %v", item.Value, item.Count)
	}

}
