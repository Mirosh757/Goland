package main

import (
	"fmt"
	"Goland/console"
	"Goland/web"
)

func main() {
	var action string
	fmt.Print("1. Запустить Веб-приложение\n2. Запустить консольное приложение\nДействие: ")
	fmt.Scan(&action)
	if action == "1"{
		web.RunServer()
	} else{
		console.Facade()
	}
}