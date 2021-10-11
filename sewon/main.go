package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strconv"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(pageNum int) { // pagination된 페이지 호출
	pageURL := baseURL + "&start=" + strconv.Itoa(pageNum*50)
	fmt.Println("Requesting: ", pageURL)

	res, err := http.Get(pageURL)
	checkErr(err)
	checkHttpStatus(res)

	defer func(Body io.ReadCloser) { // error handling for io.ReadCloser
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchJobs := doc.Find(".tapItem")

	searchJobs.Each(func(i int, selection *goquery.Selection) {
		id, _ := selection.Attr("data-jk")
		title := selection.Find("h2>span").Text()
		location := selection.Find(".companyLocation").Text()
		fmt.Println(id, title, location)
	})
}

func getPages() int { // pagination 수 반환
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkHttpStatus(res)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, selection *goquery.Selection) {
		// For each item found, count pages
		pages = selection.Find("a").Length() // <a> 태그 찾기
	})

	return pages
}

func checkErr(err error) { // 에러 처리
	if err != nil {
		log.Fatalln(err)
	}
}

func checkHttpStatus(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
