package main

import "fmt"

// byte -> uint8
// int -> int32 or int64
// unint -> uint32 or uint64

func main() {
	var i int8 = 20
	var f float32 = 3.4
	fmt.Println(float32(i) + f)

	fmt.Println(i + int8(f+1.2))

}
