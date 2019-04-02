package main

import "fmt"

func main() {
	ch := GetReciver()

	//panic because you cannot send value to receive-only  type of channel
	//ch <- 5
	for d := range ch {
		fmt.Println(d)
	}

	ch2 := GetPrinter()

	//panic because you cannot read value from send-only  type of channel
	//val := <-ch2
	for i := 10; i < 20; i++ {
		ch2 <- i
	}

}

func GetReciver() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func GetPrinter() chan<- int {
	ch := make(chan int)
	go func() {
		for d := range ch {
			fmt.Println(d)
		}
	}()

	return ch
}
