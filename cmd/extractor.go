package cmd

import (
	"github.com/PuerkitoBio/goquery"
)

func Extract(docs []*goquery.Document) ([]string, error) {
	var texts []string
	for _, doc := range docs {
		text := extract(doc)
		texts = append(texts, text)
	}
	return texts, nil
}

func extract(doc *goquery.Document) string {
	doc.Find("script").Each(func(i int, el *goquery.Selection) {
		el.Remove()
	})
	doc.Find("style").Each(func(i int, el *goquery.Selection) {
		el.Remove()
	})
	doc.Find("noscript").Each(func(i int, el *goquery.Selection) {
		el.Remove()
	})

	return doc.Find("body").Text()
}
