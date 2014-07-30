package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

/*
Images

image パッケージは、以下の Image インターフェースを定義しています：

package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
（詳細は、 このドキュメント を参照してください。）

また、 color.Color と color.Model は共にインターフェースです。 しかし、定義済みの color.RGBA と color.RGBAModel を使うことによって、そのインターフェースを無視します。
*/
