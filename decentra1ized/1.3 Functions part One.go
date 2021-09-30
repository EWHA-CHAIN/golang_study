package main

import "fmt"
import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	name := "nico"
	name = "lynn"
	fmt.Println(name)
	totalLenght, _ := lenAndUpper("nico")
	fmt.Println(totalLenght)
}