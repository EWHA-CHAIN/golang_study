package main

import (
	"fmt"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println(multiply(2, 2))
	// totalLenght, upperName := lenAndUpper("nico")
	totalLenght, _ := lenAndUpper("nico")
	fmt.Println(totalLenght)

	repeatMe("nico", "lynn", "kim", "han")

}
