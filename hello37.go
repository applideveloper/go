package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

/*
Maps

map はキーと値とを関連付けます。

mapは、使用する前に make ( new じゃない)で作成する必要があります。 nil のmapは空なので、要素を割り当てることはできません。
*/
