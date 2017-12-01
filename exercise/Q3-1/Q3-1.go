package main

import "fmt"

func main() {

	number := 0
	str := ""
	for number < 100 {
		str = str + "A"
		fmt.Printf("%s\n", str)
		number = number + len(str)
	}
}
