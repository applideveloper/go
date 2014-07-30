package main

/*
Imports
例のコードは、括弧でインポートをグループ化し、要素化した( factored )インポートステートメントとしています。 もちろん、複数のインポートステートメントを用いて以下のように書くこともできます：
*/

import (
    "fmt"
    "math"
)

func main() {
    fmt.Printf("Now you have %g problems.",
        math.Nextafter(2, 3))
}
