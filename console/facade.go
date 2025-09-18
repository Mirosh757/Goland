package console

import "fmt"

func Facade(){
	fmt.Println("Меню управления списком животных")
	fmt.Println("1. Добавить новое животное")
	fmt.Println("2. Посмотреть весь список животных")
	fmt.Println("3. Найти животное по id")
	fmt.Println("4. Обновить данные о животном")
	fmt.Println("5. Удалить информацию об одном животном")
	fmt.Println("Для выхода укажите любой другой символ")
	for {
		var action string
		fmt.Print("Действие: ")
		fmt.Scan(&action)
		switch (action){
			case "1":
				createAnimals()
			case "2":
				readAllAnimals()
			case "3":
				readOneAnimals()
			case "4":
				updateAnimals()
			case "5":
				deleteOneAnimals()
		}
	}
}