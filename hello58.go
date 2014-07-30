package main

import (
	"fmt"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (str String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintf(w, "%s", str)
}

func (str *Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintf(w, "%s%s%s", str.Greeting, str.Punct, str.Who)
}

func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	http.ListenAndServe("localhost:4000", nil)
}

/*
Exercise: HTTP Handlers

以下の型を実装し、それらにServeHTTPメソッドを定義してみてください。 自分のWebサーバで、特定のパスを処理（ハンドル）できるように登録してみてください。

type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}
例えば、以下のようなハンドラを用いて登録できるようにする必要があります。

http.Handle("/string", String("I'm a frayed knot."))
http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
*/
