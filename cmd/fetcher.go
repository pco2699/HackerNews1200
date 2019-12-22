package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	TopStories = "https://hacker-news.firebaseio.com/v0/topstories.json"
	Story      = "https://hacker-news.firebaseio.com/v0/item/{{number}}.json"
)

func Fetch() ([]*goquery.Document, error) {
	stories, err := fetchTopStories()
	if err != nil {
		return nil, err
	}
	articles, err := fetchArticles(stories)
	if err != nil {
		return nil, err
	}
	docs := fetchDocuments(articles)

	return docs, nil
}

func fetchDocuments(articles []map[string]interface{}) []*goquery.Document {
	var docs []*goquery.Document
	for _, article := range articles {
		if article["url"] != nil {
			doc, err := fetchHtml(article["url"].(string))
			if err != nil {
				continue
			}
			docs = append(docs, doc)
		}
	}

	return docs
}

func fetchHtml(url string) (*goquery.Document, error) {
	if strings.Contains(url, "pdf") || strings.Contains(url, "PDF") {
		return nil, fmt.Errorf("skip the pdf")
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}
	return doc, nil
}

func fetchTopStories() ([]int, error) {
	if stories, err := fetch(TopStories); err == nil {
		return stories, nil
	} else {
		return nil, err
	}
}

func fetch(url string) ([]int, error) {
	// httpで該当のURLへアクセス
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// レスポンスのbodyをすべてbytes[]型で取得
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// json.Unmarshalでjsonをint[]のarrayにマッピングする
	var array []int
	err = json.Unmarshal(bytes, &array)
	if err != nil {
		return nil, err
	}

	return array, nil
}

func fetchArticles(stories []int) ([]map[string]interface{}, error) {
	var articles []map[string]interface{}
	for _, s := range stories {
		article, err := fetchMap(replaceUrl(Story, s))
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func fetchMap(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var array interface{}
	err = json.Unmarshal(bytes, &array)
	if err != nil {
		return nil, err
	}

	return array.(map[string]interface{}), nil
}

func replaceUrl(url string, number int) string {
	return strings.Replace(url, "{{number}}", strconv.Itoa(number), 1)
}
