package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func search(str string, urls []string, contentType string) ([]string, error) {
	res := make([]string, 0, len(urls))

	client := http.Client{}

	for _, url := range urls {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		if strings.HasPrefix(resp.Header.Get("Content-Type"), contentType) {
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		if bytes.Contains(body, []byte(str)) {
			res = append(res, url)
		}
	}

	return res, nil
}

func main() {
	urls := []string{
		"https://google.ru",
		"https://yandex.ru",
		"https://2ip.ru",
		"https://play.golang.org",
	}
	searchString := "golang"

	result, err := search(searchString, urls, "text/plain")
	if err != nil {
		panic(err)
	}

	for _, url := range result {
		fmt.Println(url)
	}
}
