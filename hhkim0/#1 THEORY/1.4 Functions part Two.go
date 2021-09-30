//#1.4 Functions part Two
package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("I'm done") // func 끝나고 실행
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return // naked return
}

func main() {
	totalLenght, uppercase := lenAndUpper("nico")
	fmt.Println(totalLenght, uppercase)

}
