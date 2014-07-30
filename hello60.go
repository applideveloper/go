package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w, h int
}

func (r *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, r.w, r.h)
}

func (r *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (r *Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := &Image{256, 256}
	pic.ShowImage(m)
}

/*
Exercise: Images

前に書いた、画像ジェネレーターを覚えてますか？ 他のものを書いてみましょう。でも今回は、データのスライスの代わりに image.Image の実装を返すようにしてみます。

Image 型を定義して、 必要なメソッド を実装し、 pic.ShowImage を呼び出してみてください。

Bounds は、 image.Rect(0, 0, w, h) のように image.Rectangle を返す必要があります。

ColorModel は、 color.RGBAModel を返す必要があります。

At は、ひとつのcolorを返す必要があります。 最後の画像ジェネレーターの値 v は、 color.RGBA{v, v, 255, 255} のものと一致します。
*/
