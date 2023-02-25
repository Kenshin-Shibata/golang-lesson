package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "mac"
}

func main() {
	switch os:= getOsName(); os {
	case "mac":
		fmt.Println("MAC!")
	case "windows":
		fmt.Println("WINDOWS!")
	case "default":
		fmt.Println("DEFAULT!")
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	}
}