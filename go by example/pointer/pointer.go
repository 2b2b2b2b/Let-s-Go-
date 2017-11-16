package main

import "fmt"

func zeroptr(ival int) {
	ival = 0
}

func zeproptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println(i)

	zeroptr(i)
	fmt.Println(i)

	zeproptr(&i)
	fmt.Println(i)
	fmt.Println("pointer:", &i)

}
