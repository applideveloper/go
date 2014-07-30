package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	p = Vertex{1, 2}  // has type Vertex
	q = &Vertex{1, 2} // has type *Vertex
	r = Vertex{X: 1}  // Y:0 is implicit
	s = Vertex{}      // X:0 and Y:0
)

func main() {
	fmt.Println(p, q, r, s)
}

/*
Struct Literals

structリテラルは、フィールドの値を列挙することによって、構造体の初期値の割り当てを示しています。

Name: 構文を使って、フィールドの一部だけを記述することができます。 （この方法での名前のフィールドの指定順序は無関係です。）

訳注：例では X: 1 として X だけを初期化しています。

特別な接頭辞 & は、新しく割り当てられたstructへのポインタを示します。
*/
