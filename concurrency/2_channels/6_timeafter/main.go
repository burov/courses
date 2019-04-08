package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		//	time.Sleep(2 * time.Second)
		ch <- "Hello World"
	}()
	select {
	case data := <-ch:
		fmt.Println(data)
	case <-time.After(time.Second):
		fmt.Println("Timeout exided")
	}
}
