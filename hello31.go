package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])

	// missing low index implies 0
	fmt.Println("p[:3] ==", p[:3])

	// missing high index implies len(s)
	fmt.Println("p[4:] ==", p[4:])
}

/*
Slicing slices

sliceは、同じ配列を参照する新しいsliceを作成することで、再sliceすることができます。

この：

s[lo:hi]
という表現では、 lo から hi-1 までの要素のsliceを導いています。 そのため、

s[lo:lo]
は、空で、

s[lo:lo+1]
は、ひとつの要素をもつことになります。
*/
