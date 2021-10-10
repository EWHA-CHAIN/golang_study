package main

//3.1, 3.6, 3.7
import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

//var errRequestFailed = errors.New("request failed")

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}

////3.2

// func main() {
// 	go sexyCount("nico")
// 	go sexyCount("flynn") //여기까지만 하면 아예 실행x..
// 	time.Sleep(time.Second * 5)
// }

// func sexyCount(person string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(person, "is sexy", i)
// 		time.Sleep(time.Second)
// 	}
// }

////3.5
// func main() {
// 	c := make(chan string)
// 	people := [5]string{"nico", "flynn", "dal", "japanguy", "larry"}
// 	for _, person := range people {
// 		go isSexy(person, c)
// 	}
// 	fmt.Println("Waiting for messages")
// 	// resultOne := <-c
// 	// resultTwo := <-c
// 	// fmt.Println("Received this message:", resultOne)
// 	// fmt.Println("Received this message:", resultTwo)
// 	for i := 0; i < len(people); i++ {
// 		fmt.Println(<-c)
// 	}
// }

// func isSexy(person string, c chan string) {
// 	time.Sleep(time.Second * 10)
// 	c <- person + " is sexy"
// }
