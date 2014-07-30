package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
)

func Walk(t *tree.Tree, c chan int) {
	if t != nil {
		_walk(t, c)
	}
	close(c)
}

func _walk(t *tree.Tree, c chan int) {
	if t != nil {
		_walk(t.Left, c)
		c <- t.Value
		_walk(t.Right, c)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for v1 := range c1 {
		v2 := <-c2
		if v1 != v2 {
			return false
		}
	}

	_, ok := <-c2
	if ok {
		return false
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

/*
Exercise: Equivalent Binary Trees

葉に保持されている値が同じ順序をもつ、多くの種類のバイナリツリーがあります。 例えば、ここに"1, 1, 2, 3, 5, 8, 13"を保持する２つのバイナリツリーがあります。


２つのバイナリツリーが同じ順序で保持しているかどうかを確認する機能は、多くの他の言語でとても複雑です。 私たちは簡単な解決法を記述するために、Goの並行性( concurrency )とチャネルを利用してみます。

この例は、型を以下のように定義する tree パッケージを利用します。

type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
*/
