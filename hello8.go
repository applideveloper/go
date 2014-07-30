package main

import "fmt"

func add(x, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(42, 13))
}

/*
Functions continued

２つ以上の関数の引数が同じ型である場合には、最後の型を残して省略することができます。

この例では、

x int, y int
を

x, y int
へ省略できます。
*/
