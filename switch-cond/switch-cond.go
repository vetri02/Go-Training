package main

import (
	"fmt"
	"os"
)

func main() {
	word := os.Args[1]
	greet := "greetings"
	switch l := len(word); word {
	case "hi":
		fmt.Println("Very formal")
		fallthrough
	case "hello":
		fmt.Println("Hi, yourself")
	case "farewell":
	case greet:
		fmt.Println("Salutations!")
	case "goodbye", "bye":
		fmt.Println("So long!")

	default:
		fmt.Println("I don't know what you said", l, "letters")
	}

	switch l := len(word); {
	case l == 5:
		fmt.Println("its 5 charaters")
	default:
		fmt.Println("I don't know what you said", l, "letters")
	}
}
