package main

import (
	"fmt"
	"time"
)

func channelA(c chan int) {
	time.Sleep(time.Second * 1)
	c <- 1
	return
}

func channelB(c chan int) {
	time.Sleep(time.Second * 2)
	c <- 2
	return
}

func main() {
	c := make(chan int)

	go channelA(c)
	go channelB(c)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c:
			fmt.Println("received 1", msg1)
		case msg2 := <-c:
			fmt.Println("received 2", msg2)
		}
	}
}
