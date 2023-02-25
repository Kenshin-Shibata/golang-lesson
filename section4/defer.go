package main

import "fmt"

func foo(){
	defer fmt.Println("world foo")

	fmt.Println("hello foo")
}

func main() {
	// deferは遅延実行
	// foo()

	// defer fmt.Println("world")

	// fmt.Println("hello")

	// スタッキングdefer => 始めに入れたものが最後に実行される success->3->2->1 の順で実行される
	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")
}