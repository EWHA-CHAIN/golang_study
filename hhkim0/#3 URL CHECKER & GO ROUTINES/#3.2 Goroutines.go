package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Go routine will be alive for as long as the program is running
	// 	  as long as the main fuction is running
	// 2. Main funtion doesn't wait for the go routines.
	go sexyCount("nico")
	go sexyCount("flynn")
	time.Sleep(time.Second * 5)
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
