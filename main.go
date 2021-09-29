package main

import (
  "fmt"
  "strings"
  // "/src/something"
)

import "fmt"

func multiply(a, b int) int {
  return a * b
}

func lenAndUpper(name string) (int, string){
  return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func lenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("I'm done")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// func canIDrink(age int) bool {
// 	if koreanAge := age + 2; koreanAge < 18 {
// 	switch koreanAge := age + 2; koreanAge {
// 	case 10:
// 		return false
// 	case 18:
// 		return true
// 	}
// 	return true
// 	return false
// }

func canIDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
  fmt.Println("Hello World!")
  // something.SayHello()

  // const name string = "nico"
  // name = "Lynn"
  //상수

  name := "nico" 
  // var name string = "nico"
  //위와 같은 뜻
  //name := false
  name = "lynn"
  fmt.Println(name)

  fmt.Println(multiply(2, 2))

  totalLength, upperName := lenAndUpper("nico")
  fmt.Println(totalLength, upperName)

  totalLenght, _ := lenAndUpper("nico")
	fmt.Println(totalLenght)

  totalLenght, up := lenAndUpper("nico")
	fmt.Println(totalLenght, up)

  result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)

  fmt.Println(canIDrink(16))

  fmt.Println(canIDrink(18))

  a := 2
	b := &a
	// a = 10
	// fmt.Println(a, *b)

	*b = 202020
	fmt.Println(a)
	names := []string{"nico", "lynn", "dal"}
	names = append(names, "flynn")
	fmt.Println(names)

  // names := []string{"nico", "lynn", "dal"}
	// names = append(names, "flynn")
	// fmt.Println(names)
	// nico := map[string]string{"name": "nico", "age": "12"}
	// for key, value := range nico {
	// 	fmt.Println(key, value)
	// }

  nico := map[string]string{"name": "nico", "age": "12"}
	for key, value := range nico {
		fmt.Println(key, value)
	}
	favFood := []string{"kimchi", "ramen"}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico.name)
}
