package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	// ここでの[2]は配列のサイズ
	var b [2]int = [2]int{100, 200}
	fmt.Println(b)
}