package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

func main() {
    fmt.Println("Welcome to the playground!")

    fmt.Println("The time is", time.Now())

    fmt.Println("And if you try to open a file:")
    // fmt.Println(os.Open("filename"))
    fmt.Println(os.Open("hello.go"))

    fmt.Println("Or access the network:")
    // fmt.Println(net.Dial("tcp", "google.com"))
    fmt.Println(net.Dial("tcp", "google.com:80"))
}

/*
The Go Playground
このツアーで実行するプログラムは、golang.org サーバ上のGo Playgroundで動いています！

このサービスはGoのプログラムをサーバで受け取り、コンパイルし、リンクし、実行し、そしてその結果を返してくれます。

ただし、playgroundで実行できるプログラムには制限があります：

playgroundは、標準ライブラリを使うことが出来ます。が、ネットワークアクセスやファイルアクセスは出来ません。ですので、playgroudのプログラムで使えるアウトプットは標準出力と標準エラーのみとなります。
playground上はいつも "2009-11-10 23:00:00 UTC" (この値の意味は、読者のために残しておきます(^^)）です。これにより、同じ出力結果を得ることが容易になります。
実行時のCPUとメモリ使用量の制限があり、プログラムはシングルスレッドで動くようになっています（ですが、多くのgoroutineを使えます）。
playgroundは、最終安定リリース版のGo環境を使っています。

次へ進みましょう！

*/
