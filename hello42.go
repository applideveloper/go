package main

import (
	"fmt"
	"math"
)

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))
}

/*
Function values

関数も値で扱えます。

つまり、関数そのものを変数に入れることもできるということです。
*/
