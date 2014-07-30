package main

import "fmt"

func add(x int, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(42, 13))
}

/*
Functions

関数は０個以上の引数を取ることができます。

この例では、 add 関数は、 int 型の２つのパラメータを取ります。

型が変数名の 後 にくることに注意してください。

（型をなぜこのように宣言するのか、についての詳細な情報は、 記事「Go's declaration syntax」 を参照してください。）
*/
