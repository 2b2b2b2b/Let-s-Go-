package main

import "fmt"

import "time"

func worder(done chan bool) {
	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Print("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worder(done)

	<-done
}
