package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var w Writer

	// os.Stdout implements Writer
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}

/*
Interfaces are satisfied implicitly

ある型に、メソッドを実装することで、インターフェースを実装します。

インターフェースを実装することを明示的に宣言する必要はありません。

この暗黙的なインターフェースは、インターフェースを定義するパッケージと、実装されているパッケージとを分離することができます。 （他に依存しません）

これはまた、明瞭なインターフェースの定義づけを促します。 なぜなら、すべての実装を見つける必要はありませんし、新しいインターフェースの名前でそれを関連付ける必要もないからです。

io パッケージは、ここで定義したインターフェース Reader と Writer を定義しています。 ですので、何もしなくて大丈夫です。

（訳注：fmt.Fprint関数の第一引数はio.Writerタイプであり、この例で定義したWriterと同じインターフェースをもっています。Goではインターフェースを実装していることを明示する必要がありませんので、このようなことが簡単に実現できます。）
*/
