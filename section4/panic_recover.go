package main

import "fmt"

func thidPartyConnectionDB() {
	panic("Unable to connect database")
}

func save() {
	defer func(){
		// 2.システムを強制終了させないようにrecoerさせる
		s := recover()
		fmt.Println(s)
	}()
	// 1.panicを起こす ※panicはGOでは非推奨
	thidPartyConnectionDB()
}

func main() {
	save()
	fmt.Println("OK?")
}