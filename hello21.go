package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

/*
if ステートメントは C言語 や Java と似ていますが、括弧 ( ) は不要で、中括弧 { } が必要です。

(もうおなじみですよね！)
*/
