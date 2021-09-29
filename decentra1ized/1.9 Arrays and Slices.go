package main
import "fmt"

func main() {
	a := 2
	b := &a
	*b = 202020
	fmt.Println(a)
	names := []string{"nico", "lynn", "dal"}
	names = append(names, "flynn")
	fmt.Println(names)
}