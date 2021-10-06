package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string) // 2. data type 명시
	people := [5]string{"nico", "flynn", "dal", "japanguy", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Print("wainting for", i)
		fmt.Println(<-c) // 1. blocking operation
		// Concurrency 동시성
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + "is sexy" // 3. arrow로 메세지 보내기
}
