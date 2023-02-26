package main

import "fmt"

type Vertex struct {
	// 小文字にしてしまうと(X -> x)外部からアクセスできないprivateになってしまう。
	// main関数からアクセスできない
	X int
	Y int
	S string
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	v.X = 100
	fmt.Println(v.X, v.Y)

	v2 := Vertex{X: 1}
	fmt.Println(v2)

	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Println(v4)

	// 宣言しただけだが、Vertexはnilではないので{0, 0}が返る
	var v5 Vertex
	fmt.Println(v5)
	*/
}