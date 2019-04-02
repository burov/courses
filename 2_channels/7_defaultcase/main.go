package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			ch <- "Some important data"
		}
	}()
	for {
		select {
		case data := <-ch:
			fmt.Println(data)
		default:
			fmt.Println("Default case")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
