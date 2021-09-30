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


func main() {
	// go의 진입점 (node.js, python과의 차이 )
	// 오로지 컴파일을 위함, 컴파일 안하려면 안해도 된다!

	fmt.Println(multiply(2,2))

	totalLength, upperName := lenAndUpper("soobin")
	fmt.Println(totalLength,upperName)

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
