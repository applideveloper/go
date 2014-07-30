package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

/*
Channels

チャネル( Channel )は、チャネルオペレータの <- を用いて値の送受信ができる直通ルートの型です。

ch <- v    // v をチャネル ch へ送る。
v := <-ch  // ch から受信し、
           // 変数を v へ割り当てる。
（データは、矢印の方向に流れます。）

マップとスライスのように、チャネルは使う前に以下のように生成する必要があります：

ch := make(chan int)
デフォルトでは、片方が準備できるまで送受信はブロックされます。 これは、明確なロックや条件変数がなくても、goroutineの同期を許します。
*/
