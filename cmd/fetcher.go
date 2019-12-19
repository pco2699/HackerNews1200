package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	TOP_STORIES = "ttps://hacker-news.firebaseio.com/v0/topstories.json"
	STORY       = " https://hacker-news.firebaseio.com/v0/item/{{number}}.json"
)

func Fetch() error {
	stories, err := fetch(TOP_STORIES)
	if err != nil {
		return err
	}
	//for _, story := range stories {
	//	fetch()
	//}

	return nil
}

func fetch(url string) ([]int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Err: %v", err)
	}
	defer resp.Body.Close()
	var arr []int

	body, err := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &arr)

	fmt.Println(arr)

	return arr, nil
}

func replaceUrl(url string, number int) string {
	return strings.Replace(url, "{{number}}", strconv.Itoa(number), 1)
}
