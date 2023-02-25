package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
		}
		if i > 5 {
			fmt.Println("break")
		}
		fmt.Println(i)
	}

	// forの省略形
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)
}