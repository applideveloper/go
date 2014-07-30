package main

import "fmt"

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    fmt.Println(split(17))
}

/*
Named results

関数はパラメータを取ることができます。Goでの関数は、ひとつだけでなく、複数の 戻り値パラメータ ( result parameters )を返すことができます。 それらには、名前を付けて変数のように扱うことができます。

もし、戻り値パラメータに名前が付けられているなら、 return ステートメントに戻り値の変数名を書く必要はありません。 return だけで構いません！
*/
