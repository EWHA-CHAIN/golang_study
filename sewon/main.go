package main // 어떤 패키지를 사용하는지 명시해줌, main.go 파일의 경우 컴파일을 위해서 필요한 것임(필수)

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
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
	defer fmt.Println("func is done") // func이 끝난 후 실행되는 코드
	return name
}

func forLoopEx(numbers ...int) int {
	total := 0
	for _, value := range numbers { // range gives index, value
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
	if koreanAge := age + 2; koreanAge < 18 { // variable expression
		return false
	}
	return true
}

func main() {
	//	name := "sewon" // const name string = "sewon" 와 동일한 의미임, Go가 type을 예측해서 찾아줌
	//	fmt.Println(name)

	//	fmt.Println(multiply(2, 3))

	//	totalLength, _ := lenAndUpper("sewon")
	//	fmt.Println(totalLength)

	//	multipleArguments("red", "orange", "yellow", "green", "blue", "purple")

	//	totalLength, name := nakedReturn("hello")
	//	fmt.Println(totalLength, name)

	//	fmt.Println(deferEx("sewon"))

	//	fmt.Println(forLoopEx(1, 2, 3, 4))
	//	forLoopExWithoutRange(1, 2, 3, 4)

	fmt.Println(ifExWithVariableExpression(20))
}
