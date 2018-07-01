package main

import (
	"fmt"
)

func addOne(a int) int {
	return a + 1
}

func addTwo(a int) int {
	return a + 2
}

func printOperation(a int, f func(int) int) {
	fmt.Println(f(a))
}

func makerAdd(b int) func(int) int {
	return func(a int) int {
		return a + b
	}
}

func makeDoubler(f func(int) int) func(int) int {
	return func(a int) int {
		b := f(a)
		return b * 2
	}
}

func main() {
	// myadd := func(a int, b int) {
	// 	fmt.Println(a + b)
	// }

	// myadd(1, 2)

	printOperation(1, addOne)
	printOperation(2, addTwo)

	add1 := makerAdd(1)
	add2 := makerAdd(1)

	fmt.Println(add1(2))
	fmt.Println(add2(20))

	doubleAdd1 := makeDoubler(add1)
	fmt.Println(doubleAdd1(5))
}
