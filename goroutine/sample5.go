package main

import (
	"fmt"
)

func looper(n int, c chan int) {

	for i := 0; i < n; i++ {
		c <- i
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go looper(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
