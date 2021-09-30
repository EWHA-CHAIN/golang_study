//#1.5 for, range, ...args
package main

import "fmt"

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers { // `range` gives index
		total += number
	}
	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
