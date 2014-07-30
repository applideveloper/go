package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

/*
Range and Close

送り手は、これ以上の送信する値がないことを示すため、チャネルを close できます。 受け手は、受信の式に２つ目のパラメータを与えることで、そのチャネルがcloseしてしまっているかどうかをテストできます。

v, ok := <-ch
もし、これ以上の受信する値がなく、チャネルが閉じているなら、 ok は、 false になります。

ループの for i := range c は、チャネルが閉じられるまで、繰り返しチャネルから値を受信します。

注意： 送り手のチャネルだけを閉じるようにしてください。受け手ではありません。 もし閉じたチャネルへの送信すると、パニック( panic )になるでしょう。

もう一つ注意： チャネルは、ファイルと同じようなものではありません。 通常は、閉じる必要はありません。 閉じるのは、これ以上来る値がないことを受け手が知らせなければならないときにだけ必要です。 例えば、 range ループを終了することなどです。
*/
