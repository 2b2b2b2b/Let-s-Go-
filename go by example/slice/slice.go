package main

import "fmt"

func main() {

	s := make([]string, 3, 5)
	fmt.Println("emp:", s)
	fmt.Println("cap", cap(s))
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("cap", cap(s))
	s = append(s, "d")

	fmt.Println("append", s)
	fmt.Println("length", len(s))
	fmt.Println("cap", cap(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println('l', l)

}
