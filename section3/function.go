package main

import "fmt"

func add(x int, y int) (int, int) {
	return x + y, x - y
}

func cal(price int, item int) (result int) {
	result = price * item
	return result
}

func convert(prince int) float64 {
	return float64(prince)
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	// 関数内関数というもの
	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)

	// ファンクションに書いてそのまま実行する
	func(x int) {
		fmt.Println("inner func", x)
	}(1)
}