package main // 어떤 패키지를 사용하는지 명시해줌, main.go 파일의 경우 컴파일을 위해서 필요한 것임(필수)
import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("request failed")

func main() {
	var results = make(map[string]string) // map 생성 및 초기화, make() : built-in func
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
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking: ", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
