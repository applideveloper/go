package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

/*
Switch with no condition

条件のないswitchは、 switch true と書いたことと同じです。

このswitchの構造は、長い長い"if-then-else"のつながりをシンプルに表現できます。
*/
