package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./lesson.go")
	if err != nil {
		log.Fatalln("Error!")
	}
	defer file.Close()
	data := make([]byte, 100)
	// ここでのerrは:=で初期化している
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("error")
	}
	fmt.Println(count, string(data))

	err = os.Chdir("test")
	if err != nil {
		log.Fatalln("Error")
	}
}