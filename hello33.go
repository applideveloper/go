package main

import "fmt"

func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}

/*
Nil slices

sliceの初期値は nil です。

nil の slice は、０の長さと容量です。

(sliceについてより詳しく学ぶには、" Slices: usage and internals " を読んでみてください)
*/
