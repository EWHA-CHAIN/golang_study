package main

import (
	"fmt"
	"net/http"
)

type requestResult struct { // http 요청 값을 저장할 구조체
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	channel := make(chan requestResult)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}
	for _, url := range urls {
		go hitURL(url, channel) // using goroutine
	}

	for i := 0; i < len(urls); i++ {
		result := <-channel
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}

}

func hitURL(url string, channel chan<- requestResult) { // chan<- : Send Only
	res, err := http.Get(url)
	status := "OK"
	if err != nil || res.StatusCode >= 400 {
		status = "FAILED"
	}
	channel <- requestResult{url: url, status: status}
}
