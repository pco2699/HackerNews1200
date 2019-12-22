package cmd

import (
	"github.com/pco2699/hackernews1200/collections"
	"gopkg.in/jdkato/prose.v2"
)

func Count(tokens []prose.Token) []collections.CounterItem {
	counter := collections.NewCounter()
	for _, token := range tokens {
		counter.AddItems(token.Text)
	}

	return counter.MostCommon(1200)
}
