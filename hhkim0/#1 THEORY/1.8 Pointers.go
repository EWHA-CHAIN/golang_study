// #1.8 Pointers!
package main

import "fmt"

func main() {
	a := 2
	b := &a // if you want to see address of anything
	// &: address
	*b = 20 // b is a pointer to a
	// if want you to see the value of the address
	// *: see through
	fmt.Println(a)
}
