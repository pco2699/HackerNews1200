package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/pco2699/hackernews1200/cmd"
)

func main() {
	fmt.Println("Fetching HackerNews API...")
	docs, err := cmd.Fetch()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Extracting HTML...")
	texts, err := cmd.Extract(docs)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Tokenize HTMLs...")
	tokens, err := cmd.Tokenize(texts)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Counting tokens...")
	items := cmd.Count(tokens)

	file, err := os.OpenFile("hackernews1200.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, item := range items {
		fmt.Fprintf(datawriter, "Text: %v Count: %v\n", item.Value, item.Count)
	}

	datawriter.Flush()
	file.Close()

}
