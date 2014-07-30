package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}

/*
Slices

sliceは、値の配列を参照し、長さ( length )も含まれています。

[]T は、 T 型の要素をもつsliceです。
*/
