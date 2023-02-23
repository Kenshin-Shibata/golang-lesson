package main

import "fmt"

/*
func init(){
	// initは一番最初に処理される特別なもの
	fmt.Print("Init!")
}
*/

func bazz() {
	fmt.Println("Buzz")
}

func main() {
	bazz()
	fmt.Println("Hello World", "TEST TEST")
}