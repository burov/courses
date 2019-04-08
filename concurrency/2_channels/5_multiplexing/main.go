package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			t := time.Now().UnixNano() % 10
			time.Sleep(time.Duration(t) * time.Second)
			ch1 <- "first"
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			t := time.Now().UnixNano() % 5
			time.Sleep(time.Duration(t) * time.Second)
			ch2 <- "second"
		}
	}()

	for {
		select {
		case x := <-ch1:
			fmt.Println(x)

		case x := <-ch2:
			fmt.Println(x)
		}
	}

}
