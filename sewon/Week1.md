## What are Building (앞으로의 학습 방향)
scrapper (웹 사이트로부터 데이터를 추출 해오는 도구)

kr.indeed.com의 데이터를 추출하는 웹 스크래퍼 생성

Go에 있는 데이터 처리 도구 사용(해당 도구들은 multi-core & 병행성 특징을 가짐)

---

## 1. THEORY
### 1) Main Package

```go
package main // 어떤 패키지를 사용하는지 명시해줌, main.go 파일의 경우 컴파일을 위해서 필요한 것임(필수)

func main() {

}
```

→ 컴파일러는 main package와 그 내부의 main function을 먼저 찾고 실행시킴

### 2) Packages and Imports

```go
//something.go
package something

import "fmt"

func SayHello() {
	fmt.Println("Hello")
}

func sayGoodbye() { // private
	fmt.Println("Goodbye")
}

//main.go
package main 

func main() {
	something.SayHello() // 대문자로 시작하는 함수만 접근 가능함
}
```
### 3) Variables and Constants

- 상수 및 변수 선언 
~~~go
package main 

import "fmt"

func main() {
	const name string = "sewon" // Go는 타입언어이므로 타입을 명시해줌, 상수 선언
    var num int = 100 // 변수 선언
    fmt.Println(num)
	fmt.Println(name)
}
~~~
→ 변수나 상수를 선언해두고 사용하지 않으면 'declared but not used compiler' 오류 발생함

- 변수 선언 축약형
~~~go
package main 

import "fmt"

func main() {
	name := "sewon" // const name string = "sewon" 와 동일한 의미임, Go가 type을 예측해서 찾아줌
	fmt.Println(name)
}
~~~

참고) Go 자료형 - [Go Type System Overview](https://go101.org/article/type-system-overview.html)

### 4) Function
- 함수 선언 형태 
~~~go
package main 

import "fmt"

func multiply(a, b int) int { //
	return a * b
}

func main() {
	fmt.Println(multiply(2, 3))
}
~~~

→ 형태 : func 함수이름(인자이름, 타입) 리턴타입 

- Multiple Value 반환

~~~go
package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) { // return multiple value 
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLength, _ := lenAndUpper("sewon") // _ : ignored value 
	fmt.Println(totalLength)
}
~~~

- Arguments 여러 개 받기 

~~~go
package main 

import (
	"fmt"
	"strings"
)

func multipleArguments(words ...string) { // 인자 무한으로 받기 
	fmt.Println(words) // [ , , ... , ] 배열 형태로 출력됨
}

func main() {
	multipleArguments("red", "orange", "yellow", "green", "blue", "purple")
}

// 출력결과
> [red orange yellow green blue purple]
~~~

- Naked Return
~~~go
package main 

import (
	"fmt"
	"strings"
)

func nakedReturn(name string) (length int, uppercase string) { // 반환값을 리턴타입과 같이 명시함
	length = len(name) // 이미 length 변수가 선언되었기 때문에 값만 할당해줌 
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, name := nakedReturn("hello")
	fmt.Println(totalLength, name)
}
~~~

- defer (함수가 끝난 뒤 실행되는 코드)
~~~go
package main 

import (
	"fmt"
	"strings"
)

func deferEx(name string) string {
	defer fmt.Println("func is done") // func이 끝난 후 실행되는 코드
	return name
}

func main() {
	fmt.Println(deferEx("sewon"))
}

// 출력 결과
> func is done
> sewon
~~~

### 5) for, range, loop
- loop
~~~go
// With range 
func forLoopEx(numbers ...int) {
	for index, value := range numbers { // range gives index, value
		fmt.Println(index, value)
	}
}

// Without range
func forLoopExWithoutRange(numbers ...int) {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}
~~~

### 6) If 조건문

~~~go
func ifExWithVariableExpression(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 { // variable expression
		return false
	}
	return true
}
~~~
→  variable expression : if-else 조건에만 사용하기 위해 변수를 생성했다는 의미로 알 수 있음 

### 7) Switch

~~~go
func switchEx(age int) bool {
	switch koreanAge := age + 2; {
	case koreanAge < 18:
		return false
	case koreanAge >= 18:
		return true
	}
	return true
}
~~~
→ if 조건문에서와 동일하게 switch문에서만 사용할 변수를 선언할 수 있음

### 8) Pointers
~~~go
func pointerEx() {
	a := 2  // var a int
	b := &a // var b *int
	fmt.Println(b, *b)
}
~~~

→ C 언어에서의 Pointer와 동일한 개념   

참고) [Pointers in Golang](https://golangdocs.com/pointers-in-golang)

### 9) Arrays and Slices

~~~go
func ArrayAndSliceEx() {
	colors := []string{"red", "yellow"}
	colors = append(colors, "orange") // append is built-in func, return updated slice
	fmt.Println(colors)
}
~~~

→ append 함수는 builtin.go에 작성되어 있는 built-in 함수임
→ update된 slice를 반환함

참고) [GO 컬렉션 - Slice](http://golang.site/go/article/13-Go-컬렉션---Slice)

### 10) Maps

~~~go
func mapEx() {
	info := map[string]string{"name": "apple", "age": "100"} // map[{type}]{type} : 순서대로 key, value의 타입
	for key, value := range info {
		fmt.Println(key, value)
	}
}
~~~

### 11) Structs

~~~go
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

// 출력 결과
> {apple 100 [kimchi ramen]}
~~~

→ C 언어의 구조체와 동일함

참고) [Go 구조체](http://golang.site/go/article/16-Go-구조체-struct)

---

## 추가 개념 정리
### Go 특징
- Go의 소스코드는 GOPATH 폴더 내부에서 작성해야함 

    → node.js처럼 package.json 같은 파일이 없음

    → go env 명령어 쳐서 나오는 GOPATH 내부에서 코드를 작성해야 함!
    
    참고) [[Go] Go 언어 시작하기 (VS Code 사용 및 설치)](https://soyoung-new-challenge.tistory.com/84)

- Go는 타입언어임

### GOROOT 와 GOPATH

- GOROOT   
    → Go SDK가 저장되어 있는 위치를 가리킴 

- GOPATH   
    → Workspace의 root를 가리킴 (default로 사용자의 home directory 내의 go 이름의 디렉토리임)
    
    → 구성 :
    - src/: location of Go source code (for example, .go, .c, .g, .s).
    - pkg/: location of compiled package code (for example, .a).
    - bin/: location of compiled executable programs built by Go.   
        
참고) [GOROOT and GOPATH](https://www.jetbrains.com/help/go/configuring-goroot-and-gopath.html)


---

## 추가 참고자료
[Go vs Java](https://www.bmc.com/blogs/go-vs-java/)

[Go의 장점과 단점](https://covenant.tistory.com/204)
