// #1.9 Arrays and Slices
package main

import "fmt"

func main() {
	// names := [5]string{"nico", "lynn", "dal"}
	// names[3] = "alala"
	// names[4] = "alala"
	// // names[5] = "alala" -> ERROR out of size

	names := []string{"nico", "lynn", "dal"}
	names = append(names, "flynn")
	fmt.Println(names)

}
