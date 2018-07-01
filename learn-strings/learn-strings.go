package main

import "fmt"

func main() {
	// var s string
	// s = `Hello World!`
	s := `Hello World!`
	s1 := ` 🏆`
	fmt.Println(s + s1)

	b := s[0]
	b1 := s[4]
	fmt.Println(b, b1, string(b), string(b1))

	a := s[0:5]
	a1 := s[7:10]
	a2 := s[:5]
	a3 := s[7:]
	fmt.Println(a, a1, a2, a3)

	fmt.Println(len(a))

	//Rune
	var r rune
	r = 127757
	fmt.Println(string(r))

	r1 := `🐙`
	fmt.Println(r1)

	//Array

	var arr [3]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)
}
