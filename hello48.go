package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z := complex128(1.0)
	for diff := complex128(1.0); cmplx.Abs(diff) > 1e-10; {
		diff = (cmplx.Pow(z, 3) - x) / (3 * cmplx.Pow(z, 2))
		z -= diff
	}
	return z
}

func main() {
	fmt.Println(Cbrt(2))
	fmt.Println(cmplx.Pow(Cbrt(2), 3))
}

/*
Advanced Exercise: Complex cube roots

complex64 と complex128 の型による、 Go言語の複素数( complex )について見てみましょう。 立方根( cube root )の場合、ニュートン法は、以下の式を繰り返すことになります：


"2"の立方根を探し、アルゴリズムが正しく動作するか確認してください。 math/cmplx パッケージに Pow 関数がありますので、結果を確認してみましょう。
*/
