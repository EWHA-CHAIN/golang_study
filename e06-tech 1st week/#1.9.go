package main

import "fmt"

func main() {
 names := [] string{"nico", "lynn", "dal", "lei", "yoo"}
 names = append(names, "flynn")
 names [3] = "lalalala"
 fmt.Println(names)
}