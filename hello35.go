package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

/*
Range continued

インデックスや値を " _ "（アンダーバー） へ代入することで破棄することができます。

もしインデックスだけが必要なら、例のコードで ", value" の部分を削除します。
*/
