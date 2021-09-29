package main

import (
	"fmt"
	"strings"
)
import "fmt"

func lenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("I'm done")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	totalLenght, up := lenAndUpper("nico")
	fmt.Println(totalLenght, up)
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}