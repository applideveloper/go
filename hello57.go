package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	http.ListenAndServe("localhost:4000", h)
}

/*
Web servers

http パッケージは、 http.Handler を実装する何かの値を用いることで、 HTTPリクエストの処理機能を提供します：

package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}
この例では、型 Hello は、 http.Handler を実装します。

"Hello!"を表示するには、 http://localhost:4000/ へアクセスしてください。

注意： この例は、ウェブベースのツアーでは実行できません。 このWebサーバを自分で動かしてみるには、 Goのインストール が必要となります。

（訳注：サンプルコードをコピーし、ローカルでコンパイルして動かしてみてください。）
*/
