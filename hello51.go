package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

/*
Methods continued

実際には、structだけではなく、パッケージ内で任意の型にもメソッドを定義できます。

しかし、他のパッケージからの型や基本型にはメソッドを定義できません。
*/
