package main

import (
	"fmt"
	"os/user"
	"time"
)

func bazz() {
	fmt.Println("Buzz")
}

func main() {
	bazz()
	fmt.Println("Hello World", time.Now())
	fmt.Println(user.Current())
}