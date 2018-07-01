package main

import (
	"fmt"
)

// func main() {
// 	addNumbers(1, 3)
// 	addNumbers(1, 33)
// 	addNumbers(1, -03)
// }

// func addNumbers(a int, b int) {
// 	fmt.Println(a + b)
// }

// func main() {
// 	a := addNumbers(1, 3)
// 	fmt.Println(a)

// 	a = addNumbers(1, 33)
// 	fmt.Println(a)

// 	a = addNumbers(1, -03)
// 	fmt.Println(a)
// }

// func addNumbers(a int, b int) int {
// 	return a + b
// }

// func divandRemainder(a int, b int) (int, int) {
// 	return a / b, a % b
// }

// func main() {
// 	div, remainder := divandRemainder(2, 3)
// 	fmt.Println(div, remainder)

// 	_, remainder = divandRemainder(20, 4)
// 	fmt.Println(remainder)

// 	div, _ = divandRemainder(21, 32)
// 	fmt.Println(div)
// }

func doubleFail(a int, arr [2]int, s string) {
	a = a * 2
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
	s = s + s
	fmt.Println("in doblefail", a, arr, s)
}

func main() {
	a := 1
	arr := [2]int{2, 4}
	s := "hello"
	doubleFail(a, arr, s)
	fmt.Println("in main", a, arr, s)
}
