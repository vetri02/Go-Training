package main

import "fmt"

func main() {

	//Slice

	s := make([]string, 0)
	fmt.Println(s)
	fmt.Println("Length of s:", len(s))
	s = append(s, "hello")
	fmt.Println(s)
	fmt.Println("Length of s:", len(s))
	fmt.Println("Length of s[0]:", s[0])
	s[0] = "goodbye"
	fmt.Println(s)
	fmt.Println("Length of s:", len(s))

	s1 := make([]string, 2)
	fmt.Println(s1)
	fmt.Println("Length of s:", len(s1[0]))
	s1 = append(s1, "hello")
	fmt.Println(s1)
	fmt.Println("Length of s:", len(s1[2]))

	s2 := []string{"a", "b", "c"}
	fmt.Println(s2)

	for k, v := range s2 {
		fmt.Println(k, v)
	}

	// Passed by reference whereas arrays copy values
	s3 := s2[0:2]
	fmt.Println(s3)
	s2[0] = "f"
	fmt.Println(s3)
	fmt.Println(s2)

	var s4 []string
	s4 = s3
	s4[1] = "dog"
	fmt.Println(s4)
	fmt.Println(s3)

	modslice(s3)
	fmt.Println(s3)

	unihello := "🍿 🌮"
	bytes := []byte(unihello)
	fmt.Println(bytes)
	runes := []rune(unihello)
	fmt.Println(runes)
	runes[1] = 'a'
	fmt.Println(runes)
	fmt.Println(bytes)
	fmt.Println(unihello)

	//Map

	m := make(map[string]int)
	m["hello"] = 300
	h := m["hello"]
	fmt.Println("h:", h)
	fmt.Println("m[a]", m["a"])

	if v, ok := m["hello"]; ok {
		fmt.Println("v", v)

	}

	m["hello"] = 20

	fmt.Println("h:", h)
	fmt.Println("m[hello]", m["hello"])

	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	fmt.Println("b in m2:", m2["b"])
	delete(m2, "b")
	fmt.Println("b in m2:", m2["b"])

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	m3 := map[string]int{
		"a": 1,
		"b": 2,
	}

	var m4 map[string]int
	fmt.Println("m4:", m4)
	fmt.Println("c in m3:", m3["c"])

	m4 = m3
	fmt.Println("m4:", m4)
	m4["c"] = 3
	fmt.Println("m4:", m4)
	fmt.Println("c in m3:", m3["c"])

	modMap(m4)
	fmt.Println("m4", m4)
	fmt.Println("m3", m3)

}

func modMap(m map[string]int) {
	m["mod"] = 4
}

func modslice(s []string) {
	s[0] = "hello"
}
