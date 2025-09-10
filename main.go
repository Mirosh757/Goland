package main

import (
	"Goland/database"
	"fmt"
)

func main() {
	var statusDatabase bool = database.ConnectionDatabase()
	if !statusDatabase{
		fmt.Println("Ошибка во время подключения к базе")
	}else{
		facade()
	}
}