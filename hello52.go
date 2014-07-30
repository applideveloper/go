package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())
}

/*
Methods with pointer receivers

メソッドは、名前付きの型( named type )や、名前付きの型へのポインタに関連付けられます。

ここまでで、２つの Abs 関数がありました。 ひとつは、 *Vertex のポインタ型( pointer type )で、 もうひとつは、 MyFloat の値型( value type )です。

ポインタのレシーバ( pointer receiver )を使う２つの理由があります。 １つ目は、各メソッド呼び出しで、値をコピーするのを避けるためです。 （最も効果的なのは、値型が大きな構造体の場合です） ２つ目は、そのメソッドは、そのレシーバが指している値を変更できるようになります。

Abs と Scale メソッドの定義を、 *Vertex の代わりに、レシーバとして Vertex を使って書き換えてみましょう。

Scale メソッドは、 v が Vertex の場合は、何の影響もありません。 Scale は`v`を変化させます。 v が値（ポインタではない）型の場合は、メソッドは Vertex のコピーを参照し、元の値を変化させることはできません。

Abs は、どちらかの方法でも動作します。 v を読み出すだけです。元の値を（ポインタを通じて）読み出し続けるか、値のコピーをするかはたいした問題ではありません。
*/
