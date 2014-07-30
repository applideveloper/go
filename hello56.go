package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	for diff := 1.0; math.Abs(diff) > 1e-10; {
		diff = ((math.Pow(z, 2) - x) / (2.0 * x))
		z = z - diff
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

/*
Exercise: Errors

Sqrt 関数を以前の演習からコピーし、 error の値を返すように修正してみてください。

Sqrt は、複素数をサポートしていないので、負の値が与えられたとき、nil以外のエラー値を返す必要があります。

新しいタイプの

type ErrNegativeSqrt float64
を作成してみてください。

そして、 error ErrNegativeSqrt(-2).String.Error() で、 "cannot Sqrt negative number: -2" を返すような

func (e ErrNegativeSqrt) Error() string
メソッドを作ります。

注意： Error メソッド中で、 fmt.Print(e) を呼び出すことは、無限ループにプログラムが陥ることでしょう。

最初に、 fmt.Print(float64(e)) として e を変換することより、これを避けることができます。 なぜでしょうか？

負の値が与えられたとき、 ErrNegativeSqrt の値を返すように Sqrt 関数を修正してみてください。
*/
