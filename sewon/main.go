package main

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type extractedJob struct {
	id          string
	title       string
	companyName string
	location    string
}

var baseURL = "https://kr.indeed.com/jobs?q=python&limit=50"

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	defer timeTrack(time.Now(), "JobScrapper")

	var jobs []extractedJob
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted : ", len(jobs))
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	// 한글 인코딩
	utf8bom := []byte{0xEF, 0xBB, 0xBF}
	_, encodingErr := file.Write(utf8bom)
	checkErr(encodingErr)

	w := csv.NewWriter(file)
	defer w.Flush() // defer : 지연실행

	headers := []string{"LINK", "TITLE", "COMPANY_NAME", "LOCATION"}

	writeErr := w.Write(headers)
	checkErr(writeErr)

	for _, job := range jobs { // _ : index, job : 요소 값
		jobSlice := []string{
			"https://kr.indeed.com/viewjob?jk=" + job.id,
			job.title,
			job.companyName,
			job.location}
		jobWriteErr := w.Write(jobSlice)
		checkErr(jobWriteErr)
	}
}

func getPage(pageNum int) []extractedJob { // pagination된 페이지 호출
	var jobs []extractedJob

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
		job := extractJob(selection)
		jobs = append(jobs, job) // slice
	})

	return jobs
}

func extractJob(selection *goquery.Selection) extractedJob {
	id, _ := selection.Attr("data-jk")
	title := cleanString(selection.Find("h2>span").Text())
	companyName := cleanString(selection.Find(".companyName").Text())
	location := cleanString(selection.Find(".companyLocation").Text())

	return extractedJob{
		id:          id,
		title:       title,
		companyName: companyName,
		location:    location}
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

func cleanString(str string) string { // ex) a:   b:   c:   -> "a:", "b:", "c:" -> a: b: c: (Join()의 결과)
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
