package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go func () {
		ch <- "Hello World"
	}()

	greeting := <- ch
	fmt.Println(greeting)
}
