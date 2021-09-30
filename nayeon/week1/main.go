package main

import (
	"fmt"
	"strings"

	"github.com/golang_study/nayeon/week1/something"
)

//1.3
func multiply(a, b int) int {
	return a * b
} //함수 인자와 리턴값의 type을 모두 지정해야 함.

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

func repeatMe(words ...string) {
	fmt.Println(words)
} //인자 여러개를 받는 경우 ...를 type 앞에 써줌

//1.4
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	//함수가 끝난 뒤에 실행하도록 함. 잘 사용하면 좋을 듯.
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
} //무엇을 리턴할지 지정 가능. return만 써도 알아서 잘 값을 반환. naked return이라고 함.

//1.5
func superAdd(numbers ...int) int {
	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// }
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }//위와 같은 내용의 코드
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

//1.6
// func canIDrink(age int) bool {
// 	if koreanAge := age + 2; koreanAge < 18 {
// 		return false
// 	}//;앞의 내용을 앞으로 뺄수도 있지만 더 깔끔하게 정리하려고 합친거임
// 	return true
// }

//1.7
func canIDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
} //else-if 줄일 때 사용하는 방법

//1.11
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	//1.1
	fmt.Println("Hello world!")
	//따로 export하지 않고 function 이름 첫글자를 대문자로 함.
	something.SayHello()
	//something.sayBye()
	//외부에서 사용할 수 없음

	//1.2
	//const name bool = "nico"
	//자료형 잘못 쓰면 에러
	//const name string = "nico"
	//name = "Lynn"
	//상수는 값 변경 불가. 에러임.
	//var name string = "nico"//변수
	name := "nico" //위와 같은 코드 축약 ver. 자료형 알아서 찾아줌. 단 함수 안에서만 가능!
	//name := false
	//처음 정해진 자료형이 bool이기 때문에 아래 줄에 오류남
	name = "lynn"
	fmt.Println(name)

	//1.3
	fmt.Println(multiply(2, 2))

	// totalLength, upperName := lenAndUpper("nico")
	// fmt.Println(totalLength, upperName)

	//하나의 리턴값만 받고 싶은 경우
	//totalLength, _ := lenAndUpper("nico")
	//fmt.Println(totalLength)

	repeatMe("nico", "lynn", "dal", "marl", "flynn")

	//1.4
	totalLength, upperName := lenAndUpper("nico")
	fmt.Println(totalLength, upperName)

	//1.5
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)

	//1.6 //1.7
	fmt.Println(canIDrink(16))

	//1.8
	a := 2
	b := a //값을 복사
	a = 10 //b에 영향 없음
	fmt.Println(a, b)

	aa := 2
	bb := &aa
	aa = 5               // a, b값이 다 5로 바뀜
	*bb = 20             // a, b값이 다 20으로 바뀜
	fmt.Println(&aa, bb) //메모리 주소 출력
	fmt.Println(aa, *bb) //값 출력

	//1.9
	//array
	names := [5]string{"nico", "lynn", "dal"}
	names[3] = "alala"
	names[4] = "alala"
	//names[5] = "alala"//out of range
	fmt.Println(names)

	//slice
	names2 := []string{"nico", "lynn", "dal"}
	names2 = append(names2, "flynn")
	fmt.Println(names2)

	//1.10
	nico := map[string]string{"name": "nico", "age": "12"} //다양한 자료형을 쓰고 싶을 경우 struct 사용
	fmt.Println(nico)
	for key, value := range nico {
		fmt.Println(key, value)
	}
	for key, _ := range nico {
		fmt.Println(key)
	}
	for _, value := range nico {
		fmt.Println(value)
	} //둘 중 하나 생략도 가능

	favFood := []string{"kimchi", "ramen"}
	//nicco := person{"nicco", 18, favFood}
	nicco := person{name: "nicco", age: 18, favFood: favFood} //꼴로 쓰는 것 추천
	fmt.Println(nicco)
	fmt.Println(nicco.name) //하나만 따로 출력도 가능
}
