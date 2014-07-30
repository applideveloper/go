package main

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Println("My favorite number is", rand.Intn(10))
}

/*
Packages

すべてのGoプログラムは、パッケージで構成されています。

プログラムの処理は main パッケージ内で始まります。

このプログラムでは、パッケージの "fmt" と "math/rand" をインポートしています。

規約で、パッケージ名はインポートパスの最後の要素と同じになります。 例えば、 "math/rand" パッケージは package rand ステートメントで始まるファイル群で構成されます。

（訳注：もしインポートパスが "code.google.com/p/go.net/websocket" だった場合は、 websocket になります）

注意: プログラムが実行されるplaygroundの環境は、いつも同じ状態です。 ですので、擬似乱数を返す rand.Intn はいつも同じ数を返します。 (数を強制的に変えるなら、乱数生成でシードを与える必要があります。rand.Seedを見てみてください。)
*/
