//#1.6 If with a Twist
package main

import "fmt"

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 { // variable expression(only for if)
		return false
	}
	return true

}

func main() {
	fmt.Println(canIDrink(16))
}
