package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	p := Vertex{1, 2}
	q := &p
	q.X = 1e9
	fmt.Println(p)
}

/*
Pointers

Go言語にはポインタがありますが、ポインタ演算はありません。

構造体のフィールドは、構造体のポインタを通してアクセスできます。このポインタを通じた間接的なアクセスで、とてもわかりやすくなります。
*/
