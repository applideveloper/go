package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

/*
Mutating Maps

map m へ要素(elem)の挿入や更新は：

m[key] = elem
要素の取り出しは：

elem = m[key]
要素の削除は：

delete(m, key)
キーが存在するかどうかは、２つの値の代入で確認します：

elem, ok = m[key]
もし、map m に key があれば、 ok は true になります。 なければ、 ok は false となり、`elem` は、mapの要素の型のゼロの値となります。

同様に、mapにkeyが存在しない場合に読み込んだ結果も、mapの要素の型のゼロの値となります。
*/
