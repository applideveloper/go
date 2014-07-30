package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}

/*
Map literals continued

もし、トップレベルの型が型名である場合は、リテラルの要素から推定できますので、その型名を省略することができます。
*/
