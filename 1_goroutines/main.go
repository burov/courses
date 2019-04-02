package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i)
	}

	for {
	}
}
