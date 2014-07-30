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
C言語 や Java のように、条件の前後を空にすることができます。
*/
