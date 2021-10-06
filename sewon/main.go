package main // 어떤 패키지를 사용하는지 명시해줌, main.go 파일의 경우 컴파일을 위해서 필요한 것임(필수)
import (
	"fmt"
	"github.com/sw-develop/learngo/golang_study/sewon/dictEx"
)

func main() {

	dictionary := dictEx.Dictionary{"first": "first word"}
	//	def1, err1 := dictionary.Search("first") // definition, err := dictionary["first"]와 동일한 의미인데, 저렇게 작성하는게 더 좋은 코드 형태임
	/*
		if errs1 != nil {
			fmt.Println(err1)
		} else {
			fmt.Println(def1)
		}

		err2 := dictionary.AddWord("first", "second word")
		if err2 != nil {
			fmt.Println(err2)
		} else {
			def2, _ := dictionary.Search("second")
			fmt.Println(def2)
		}
	*/

	baseWord := "fourth"
	dictionary.AddWord(baseWord, "fourth word")
	err3 := dictionary.UpdateWord(baseWord, "update next word")
	if err3 != nil {
		fmt.Println(err3)
	} else {
		value, _ := dictionary.Search(baseWord)
		fmt.Println(value)
	}

	dictionary.DeleteWord(baseWord)
	value, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}
