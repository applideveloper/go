package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i, j := 1, 1
	return func() int {
		i, j = j, i+j
		return i
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

/*
Exercise: Fibonacci closure

関数で面白いものを見てみましょう。

fibonacci 関数を実装しましょう。この関数は、連続するフィボナッチ( fibonacci )数を返す関数（クロージャ）を返します。
*/
