package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

/*
If with a short statement

if ステートメントは、 for のように、実行のための短いステートメントを条件の前に書くことができます。

ここで宣言された変数は、 if のスコープだけで有効です。

(ためしに最後の return 文で、 v を使ってみてください。 使えなかったでしょ？)
*/
