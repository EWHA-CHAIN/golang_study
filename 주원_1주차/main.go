package main

// 내가 작성할 패키지의 이름

import (
	"fmt"
	"math/big"
	"std/github.com/juwon/learngo/something"
	"strings"
)

func multiply(a int,b int) int{
	// a,b int -> a,b int로 인식
	return a*b
}

func lenAndUpper(name string)(int,string){
	return len(name),strings.ToUpper(name)
}
//go는 multiple return value의 특성 

func lenAndUpper_1(name string)(length int,uppercase string){
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words string) string{
	fmt.Println(words)
	return
}

func multipleArguments(words ...string) {
	fmt.Println(words) // [ , , ... , ] 배열 형태로 출력됨
}

func nakedReturn(name string) (length int, uppercase string) { // 반환값을 리턴타입과 같이 명시함
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func deferEx(name string) string {
	defer fmt.Println("func is done") 
	return name
}

func forLoopEx(numbers ...int) int {
	total := 0
	for _, value := range numbers { 
		total += value
	}
	return total
}

func forLoopExWithoutRange(numbers ...int) {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}

func ifExWithVariableExpression(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 { 
		return false
	}
	return true
}

func switchEx(age int) bool {
	switch koreanAge := age + 2; {
	case koreanAge < 18:
		return false
	case koreanAge >= 18:
		return true
	}
	return true
}

func pointerEx() {
	a := 2  // var a int
	b := &a // var b *int
	fmt.Println(b, *b)
}

func ArrayAndSliceEx() {
	colors := []string{"red", "yellow"}
	colors = append(colors, "orange") 
	fmt.Println(colors)
}

func mapEx() {
	info := map[string]string{"name": "apple", "age": "100"} 
	for key, value := range info {
		fmt.Println(key, value)
	}
}

type person struct {
	name         string
	age          int
	favoriteFood []string
}

func structEx() {
	favoriteFood := []string{"kimchi", "ramen"}
	info := person{name: "apple", age: 100, favoriteFood: favoriteFood}

	fmt.Println(info)
}
func main() {
	// go의 진입점 (node.js, python과의 차이 )
	// 오로지 컴파일을 위함, 컴파일 안하려면 안해도 된다!

	fmt.Println(multiply(2,2))

	totalLength, upperName := lenAndUpper("soobin")
	fmt.Println(totalLength,upperName)
	// totalLength, _ := lenAndUpper ("soobin") -> '_'는 value 값 무시
	
	length, uppercase := lenAndUpper_1(4,"soobin")
	fmt.Println(length,uppercase)
	
	repeatMe("soobin, yeonjun, hyuningkai")
	
	
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)

	fmt.Println(canIDrink(16))


	a := 2
	b := a 
	a = 10 
	fmt.Println(a, b)

	aa := 2
	bb := &aa
	aa = 5              
	*bb = 20            
	fmt.Println(&aa, bb) 
	fmt.Println(aa, *bb) 

	
	names := [5]string{"nico", "lynn", "dal"}
	names[3] = "alala"
	names[4] = "alala"
	//names[5] = "alala"//out of range
	fmt.Println(names)


	names2 := []string{"nico", "lynn", "dal"}
	names2 = append(names2, "flynn")
	fmt.Println(names2)

	
	nico := map[string]string{"name": "nico", "age": "12"} 
	fmt.Println(nico)
	for key, value := range nico {
		fmt.Println(key, value)
	}
	for key, _ := range nico {
		fmt.Println(key)
	}
	for _, value := range nico {
		fmt.Println(value)
	} 

	favFood := []string{"kimchi", "ramen"}
	//nicco := person{"nicco", 18, favFood}
	nicco := person{name: "nicco", age: 18, favFood: favFood} 
	fmt.Println(nicco)
	fmt.Println(nicco.name)
	
	fmt.Println("Hello World!")
	//go의 경우에는 함수를 import 하고 싶다면 function을 대문자로 시작하면 된다.
	//fmt=formating / vscode에서는 자동 import 되는 특성

	something.SayHello()

	// 상수는 const 로 선언 

	var name string = "soobin"
	// name := "soobin" 이라고도 쓸 수 있다. -> go가 타입 찾아줌
	// 첫번째 값의 타입에 따라 값을 변경해야 된다.
	// 축약형: func 안에서만 가능, 변수에만 가능  
	name = "yeonjun"
	//go는 타입 명시해야 된다.
	fmt.Println(name)
}
