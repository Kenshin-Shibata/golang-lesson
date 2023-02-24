package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello" + "World")
	fmt.Println(string("Hello World"[0]))

	var s string = "Hello World"
	// 文字列sの中身を指定したものと入れ替える
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)

	// 指定した文字列が含まれているかを判断する
	fmt.Println(strings.Contains(s, "World"))
}