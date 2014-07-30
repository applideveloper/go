package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

/*
Making slices

sliceは、 make 関数で生成します。 これは、ゼロに初期化した配列をメモリに割り当て、その配列を参照したsliceを返す働きをします：

a := make([]int, 5)  // len(a)=5
sliceは、長さと容量( capacity )を持っています。 sliceの容量は、sliceが基礎となる配列で拡大できる最大の長さです。

容量を指定するためには、 make の３番目の引数に渡します：

b := make([]int, 0, 5) // len(b)=0, cap(b)=5
sliceは　再スライス　( re-slicing )によって拡大縮小させることができます（最大容量まで）：

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
*/
