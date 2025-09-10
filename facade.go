package main

import "fmt"

func facade(){
	fmt.Println("Меню управления списком животных")
	fmt.Println("1. Добавить новое животное")
	fmt.Println("2. Посмотреть весь список животных")
	fmt.Println("3. Найти животное по id")
	fmt.Println("4. Обновить данные о животном")
	fmt.Println("5. Удалить информацию об одном животном")
	fmt.Println("6. Удалить список животных")
	fmt.Println("Для выхода укажите любой другой символ")
	for {
		var action string
		fmt.Print("Действие: ")
		fmt.Scan(&action)
		switch (action){
			case "1":
				readAllAnimals()
		}
	}
}