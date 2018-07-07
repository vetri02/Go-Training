package main

import (
	"fmt"
)

func doStuff(i interface{}) {
	switch i := i.(type) {
	case int:
		fmt.Println("Double i is", i+i)
	case string:
		fmt.Println("i is", len(i), "characters long")
	default:
		fmt.Println("I don't know what is this")
	}
}

func main() {
	doStuff(10)
	doStuff("dfas")
	doStuff(true)
}
