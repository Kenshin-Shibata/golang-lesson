package main

import "fmt"

func one(x *int) {
	// xはnのアドレスなので、その中身に1を入れてくださいという意味
	*x = 1
}

func main() {
	var n int = 100
	one(&n) //nのアドレスを渡す
	fmt.Println(n) // nのアドレスに1が入ったので1が出力される

	/*
	fmt.Println(n) //100

	fmt.Println(&n)

	var p *int = &n //pというポインタ型の箱の中にnのアドレスを入れているという処理
	fmt.Println(p) //なので、pに入れられたアドレスが出力される

	fmt.Println(*p) //pに入っているアドレスに何が入っているかを参照する
	*/
}