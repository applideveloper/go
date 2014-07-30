package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	//	runtime.GOMAXPROCS(2)
	go say("world") //新しいGoroutinesを実行する。
	say("hello")    //現在のGoroutines実行
}
