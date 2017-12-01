package main

import "fmt"

func main() {
	index := 0
Here:
	fmt.Println(index)

	if index < 10 {
		index++
		goto Here
	}
}
