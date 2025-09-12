package main

import (
	"database/sql"
	"fmt"
	_ "strings"
)

type animal struct {
	id           int
	name         string
	type_animals string
	breed        string
	age          int
	gender       string
	color        string
}

func createAnimals() {

}

func readAllAnimals() {
	db, err := sql.Open("sqlite", "./database/animals.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM animals")
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	animals := []animal{}

	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.name, &a.type_animals, &a.breed, &a.age, &a.gender, &a.color)
		if err != nil {
			fmt.Println(err)
			continue
		}
		animals = append(animals, a)
	}
	for _, a := range animals {
		fmt.Println(a.id, a.name, a.type_animals, a.breed, a.age, a.gender, a.color)
	}
}

func readOneAnimals() {
	db, err := sql.Open("sqlite", "./database/animals.db")
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	var id string
	fmt.Print("Укажите id: ")
	fmt.Scan(&id)

	rows, err := db.Query("SELECT * FROM animals WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	animals := animal{}

	if rows.Next() {
		err := rows.Scan(&animals.id, &animals.name, &animals.type_animals, &animals.breed, &animals.age, &animals.gender, &animals.color)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(animals.id, animals.name, animals.type_animals, animals.breed, animals.age, animals.gender, animals.color)
		}
	} else {
		fmt.Println("Введенного id не существует в базе")
	}
}

func updateAnimals() {

}

func deleteOneAnimals() {
	db, err := sql.Open("sqlite", "./database/animals.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var id string
	fmt.Print("Укажите id: ")
	fmt.Scan(&id)

	statement, err := db.Prepare("DELETE FROM animals WHERE id = ?")
	if err != nil {
		fmt.Println("Ошибка подготовки запроса:", err)
		return
	}
	defer statement.Close()

	result, err := statement.Exec(id)
	if err != nil {
		fmt.Println("Ошибка выполнения:", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Ошибка получения результата:", err)
		return
	}
	
	fmt.Printf("Удалено записей: %d\n", rowsAffected)
}

func deleteManyAnimals() {

}
