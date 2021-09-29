package main

import "fmt"

type person struct {
  name string
  age int
  favFood []string
}

func main() {
  favFood := []string{"meat", "coke"}
  nico := person{name: "nico", age: 19, favFood: favFood}
  fmt.Println(nico.name)
}