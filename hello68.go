package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

/*
Default Selection

select の中で、どの case にも一致しないのであれば、 default の case が実行されます。

ブロックしないで送受信するなら、 default の case を使ってください：

select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}
*/
