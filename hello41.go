package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	for _, v := range strings.Fields(s) {
		ret[v] += 1
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}

/*
Exercise: Maps

WordCount 関数を実装してみましょう。文字列 s の “word” の数のmapを戻す必要があります。 wc.Test 関数は、引数に渡した関数に対しテストスイートを実行し、成功か失敗かを結果に表示します。

strings.Fields で、何かヒントを得ることができるはずです。
*/
