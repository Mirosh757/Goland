package main

import (
	"fmt"
	"Goland/console"
	"Goland/web"
	 "Goland/database"
)

func main() {
	var dbUp string
	fmt.Print("Обновить базу данных?(Y - да, N - нет): ")
	fmt.Scan(&dbUp)
	if dbUp == "Y"{
		database.ConnectionDatabase()
	}
	var action string
	fmt.Print("1. Запустить Веб-приложение\n2. Запустить консольное приложение\nДействие: ")
	fmt.Scan(&action)
	if action == "1"{
		web.RunServer()
	} else{
		console.Facade()
	}
}