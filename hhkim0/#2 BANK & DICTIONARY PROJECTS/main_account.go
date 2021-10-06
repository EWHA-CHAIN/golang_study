package main

import (
	"fmt"

	"github.com/hhkim0/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())

	err := account.Withdraw(20)
	if err != nil { // Go에는 try-catch 가 없어.
		// log.Fatalln(err) // 에러나면 종료
		fmt.Println(err)
	}
	// fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}
