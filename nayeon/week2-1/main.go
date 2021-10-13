package main

import (
	"fmt"

	"github.com/golang_study/nayeon/week2-1/mydict"
)

////2.0-3
// func main() {
// 	account := accounts.NewAccount("nico")
// 	account.Deposit(10)
// 	err := account.Withdraw(20)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(account)
// }

//2.4-2.6
func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }
	//word := "hello"
	// definition := "Greeting"
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// hello, _ := dictionary.Search(word)
	// fmt.Println("found", word, "definition:", hello)
	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	// err := dictionary.Update(baseWord, "Second")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	//err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
	//fmt.Println(word)
}
