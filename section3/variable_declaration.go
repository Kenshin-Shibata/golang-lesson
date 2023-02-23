package main

import "fmt"

// func main() {
// 	var i int = 1
// 	var f64 float64 = 1.2
// 	var s string = "test"
// 	var t bool = true
// 	var f bool = false
// 	fmt.Println(i, f64, s, t, f)
// }

/*
func main() {
	// 宣言だけして値を入れない場合は初期値が入る
	var (
		i int
		f64 float64
		s string
		t, f bool
	)
	fmt.Println(i, f64, s, t, f)
}
*/

func main() {
	// 宣言だけして値を入れない場合は初期値が入る
	var (
		i int = 1
		f64 float64 = 1.2
		s string = "test"
		t, f bool = true, false
	)
	fmt.Println(i, f64, s, t, f)

	// ↓ ショートの宣言方法(ただし、関数外では使えない)
	xi := 1
	xf64  := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
}