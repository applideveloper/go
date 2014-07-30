package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

/*
For is Go's "while"

セミコロンを抜くこともできます。つまり、C言語での while は、Goでは for だけを使います。
*/
