package main
import "fmt"

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return true
	return false
}

func main() {
	fmt.Println(canIDrink(16))
	fmt.Println(canIDrink(18))
}