package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

/*
Select

select ステートメントは、goroutineを複数の通信操作で待たせます。

select は、それの case の条件で、いずれかを実行できるようになるまでブロックし、そして、条件が一致した case を実行します。 もし、複数が一致した場合、 case はランダムに選ばれます。
*/
