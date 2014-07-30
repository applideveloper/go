package main

import "fmt"

func main() {
	c := make(chan int, 2) //2を1に修正するとエラーが発生します。2を3に修正すると正常に実行されます。
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
