package main

import (
	"fmt"

	"github.com/hhkim0/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")

	// UPDATE
	// err := dictionary.Update(baseWord, "Second")		// update 성공
	// err := dictionary.Update("asdfasdf", "Second")	// update 실패
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// word, _ := dictionary.Search(baseWord)

	// DELETE
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
}

// # 2.4
func dict1() {
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}

// # 2.5
func dict2() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, err := dictionary.Search(word)
	fmt.Println("found", word, "definition:", hello)
	err2 := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err2)
	}
}
