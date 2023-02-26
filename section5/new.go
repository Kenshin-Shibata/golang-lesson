package main

import "fmt"

func main() {
	// 何も入れない状態でメモリの領域を確保する
	var p *int = new(int)
	fmt.Println(p) // 確保されたアドレスが返る
	fmt.Println(*p)
	*p++
	fmt.Println(*p)
}