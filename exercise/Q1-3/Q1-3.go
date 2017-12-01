package main

import "fmt"

func main() {
	list := [...]int{1, 2, 3, 4, 5}
	//list := [5]int{1,2,3,4,5}
	for index, value := range list {
		fmt.Printf("index:%d   ", index)
		fmt.Printf("value:%d\n", value)
	}
}
