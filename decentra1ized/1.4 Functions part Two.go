package main
import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
func lenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("I'm done")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLenght, _ := lenAndUpper("nico")
	fmt.Println(totalLenght)
	totalLenght, up := lenAndUpper("nico")
	fmt.Println(totalLenght, up)
}