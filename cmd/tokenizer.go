package cmd

import (
	"gopkg.in/jdkato/prose.v2"
)

func Tokenize(texts []string) ([]prose.Token, error) {
	var tokens []prose.Token
	for _, text := range texts {
		if t, err := tokenize(text); err == nil {
			tokens = append(tokens, t...)
		} else {
			return nil, err
		}
	}
	return tokens, nil
}

func tokenize(text string) ([]prose.Token, error) {
	doc, err := prose.NewDocument(text)
	if err != nil {
		return nil, err
	}
	return doc.Tokens(), nil
}
