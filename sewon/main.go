package main // 어떤 패키지를 사용하는지 명시해줌, main.go 파일의 경우 컴파일을 위해서 필요한 것임(필수)
import (
	"fmt"
	"github.com/sw-develop/learngo/golang_study/sewon/accounts"
)

func main() {
	account := accounts.NewAccount("meme") // 변수 선언 축약형 - var account *account.Account = account.NewAccount(" ")와 동일함
	fmt.Println(*account)

	account.Deposit(10)
	fmt.Println(account.GetBalance())

	err := account.WithDraw(20)
	if err != nil { // error handling in Go (직접 작성해줘야 함)
		fmt.Println(err)
	}
	fmt.Println(account.GetBalance(), account.GetOwner())

	fmt.Println(account)
}
