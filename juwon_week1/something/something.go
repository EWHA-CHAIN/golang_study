package something

import "fmt"

func sayBye() {
	//private function
	fmt.Println("Bye")
}

func SayHello() {
	//대문자 시작-> export 가능
	fmt.Println("Hello")
}
