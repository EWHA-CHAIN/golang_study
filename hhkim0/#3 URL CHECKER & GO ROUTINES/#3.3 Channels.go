package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy_0(person, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
	// fmt.Println(<-c)	// deadlock!
}
func isSexy_0(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}
