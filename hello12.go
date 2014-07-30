package main

import "fmt"

var i, j int = 1, 2
var c, python, java = true, false, "no!"

func main() {
    fmt.Println(i, j, c, python, java)
}

/*
Variables with initializers

var 宣言では、変数ひとつひとつに初期化子( initializer )を与えることができます。

もし初期化子が指定されている場合、型を省略できます。 その変数は初期化子の型になります。
*/
