package main

import "fmt"

import "strings"

func lenAndUpper(name string) (length int, uppercase string) {
  defer fmt.Println("I'm done")
  length = len(name)
  uppercase = strings.ToUpper(name)
  return 
}

func main() {
  totalLength,up := lenAndUpper("eunsoo")
  fmt.Println(totalLength, up)
}