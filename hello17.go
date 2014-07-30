package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

/*
For

Goは、 for ループだけを繰り返し文として使います。 Goには while 文はありません！

基本的には、C言語 や Java と同じですが、括弧 ( ) は不要で（付けてはいけません）、中括弧 { } は必要です。
*/
