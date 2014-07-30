package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
}

/*
Methods

Goには、クラス( class )のしくみはありませんが、struct型にメソッド( method )を定義できます。

メソッドレシーバ ( method receiver )は、 func キーワードとメソッド名の間で、それ自身の引数リストで表現します。 （訳注：この例では、 (v *Vertex) のことを言っている）
*/
