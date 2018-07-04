package main

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

func rot13(in rune) rune {
	if in >= 'A' && in <= 'Z' {
		return ((((in - 'A') + 13) % 26) + 'A')
	}
	if in >= 'a' && in <= 'z' {
		return ((((in - 'a') + 13) % 26) + 'a')
	}

	return in
}

func main() {
	s := "This is a string 👨‍🔬"
	s2 := strings.Map(rot13, s)
	fmt.Println(s2)
	s3 := strings.Map(rot13, s2)
	fmt.Println(s3)

	emo := "👨‍🔬 🕵️"
	fmt.Println(emo)
	fmt.Println(len(emo))
	fmt.Println(utf8.RuneCountInString(emo))

	t := time.Now()
	nanos := t.UnixNano()
	fmt.Println(t)
	fmt.Println(nanos)
}
