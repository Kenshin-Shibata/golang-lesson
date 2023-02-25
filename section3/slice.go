package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])

	var board = [][]int{
		{0,1,2},
		{3,4,5},
		{6,7,8},
	}

	fmt.Println(board)
}