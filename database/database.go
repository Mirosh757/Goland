package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// var database *sql.DB
var statusDatabase bool = false

func ConnectionDatabase() bool {
	database, err := sql.Open("sqlite", "./database/animals.db")
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		return statusDatabase
	}
	defer database.Close()

	// Проверка подключения
	err = database.Ping()
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		return statusDatabase
	}
	statusDatabase = true
	fmt.Println("Успешное подключение к базе данных")
	createTable(database)
	return statusDatabase
}

func createTable(database *sql.DB) bool {
	createSQL := `
		CREATE TABLE animals(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(100) NOT NULL,
		type_animals VARCHAR(20) NOT NULL,
		breed VARCHAR(120) NOT NULL,
		age INTEGER,
		gender VARCHAR(7),
		color VARCHAR(30))
	`
	statement, err := database.Prepare(createSQL)
	if err != nil {
		fmt.Println("Неверный запрос:", err)
		return false
	}

	_, err = statement.Exec()
	if err != nil {
		fmt.Println("Ошибка выполнения запроса:", err)
		return false
	}
	fmt.Println("Таблица успешно создана")
	insertTable(database)
	return true
}

func insertTable(database *sql.DB){
	insertSQL := `
    INSERT INTO animals (name, type_animals, breed, age, gender, color) 
    VALUES 
        ('Барсик', 'cat', 'британец', 3, 'male', 'серый'),
        ('Муся', 'cat', 'сиамская', 2, 'female', 'белый'),
        ('Шарик', 'dog', 'овчарка', 5, 'male', 'черный'),
        ('Рекс', 'dog', 'дворняжка', 4, 'male', 'рыжий')
    `
	result, err := database.Exec(insertSQL)
	if err != nil{
		fmt.Println("Неправильный запрос на вставку")
	} else{
		rowAffected, _ := result.RowsAffected()
		fmt.Printf("Добавлено %d строк\n", rowAffected)
	}
}
