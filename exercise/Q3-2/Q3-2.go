package main

import "fmt"
import "unicode/utf8"

func main() {
	str := "dsjkdshdjsdh....js"
	fmt.Printf("String %s\nLength: %d, Runes: %d\n", str, len([]byte(str)), utf8.RuneCount([]byte(str)))
	fmt.Println([]byte(str))
}
