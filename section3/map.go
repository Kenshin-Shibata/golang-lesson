package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)

	// 配列の中に"apple"が入っているか確認する -> あればtrue, なければfalse
	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	// 初期化 メモリ上に空のmapを作ってそちらに入れていく
	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	var s []int
	if s == nil{
		fmt.Println("Nil")
	}
}