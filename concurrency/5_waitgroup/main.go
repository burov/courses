package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		wg    = &sync.WaitGroup{}
		sum   = 0
		rChan = make(chan int, 9)
	)

	digits := getDigits()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(ind int) {
			sqrtDigits(rChan, digits[ind:ind+10]...)
			wg.Done()
		}(i)
	}

	wg.Wait()
	close(rChan)

	for d := range rChan {
		sum += d
	}

	fmt.Println(sum)
}

func sqrtDigits(rch chan<- int, digits ...int) {
	for _, d := range digits {
		rch <- d * d
	}
}

func getDigits() []int {
	s := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		s = append(s, 2)
	}

	return s
}
