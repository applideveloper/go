package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

/*
Goroutines

goroutine (ゴルーチン)は、Goのランタイムに管理される軽量なスレッドです。

go f(x, y, z)
と書けば、 f は新しいgoroutine上で実行されます。

f(x, y, z)
f 、`x 、 y`、 z の評価は、現在のgoroutineで発生し、 f の実行は、新しいgoroutineで発生します。

goroutineは、同じアドレス空間で実行されるため、共有メモリへのアクセスは、きちんと同期する必要があります。 sync パッケージは役に立つ方法を提供していますが、他の方法があるので、Goではあまり必要ありません。 （次のスライドをみてください。）
*/
