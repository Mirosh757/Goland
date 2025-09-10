package main

import (
	"fmt"
	//"strings"
	"database/sql"
)

type animal struct{
	id int
	name string
	type_animals string
	breed string
	age int
	gender string
	color string
}

func createAnimals(){

}

func readAllAnimals(){
	db, err := sql.Open("sqlite", "./database/animals.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

	rows, err := db.Query("SELECT * FROM animals")
	if err != nil{
		panic(err)
	}
	defer db.Close()
	animals := []animal{}
	
	for rows.Next(){
		a := animal{}
		err := rows.Scan(&a.id, &a.name, &a.type_animals, &a.breed, &a.age, &a.gender, &a.color)
		if err != nil{
			fmt.Println(err)
			continue
		}
		animals = append(animals, a)
	}
	for _, a := range animals {
		fmt.Println(a.id, a.name, a.type_animals, a.breed, a.age, a.gender, a.color)
	}
}

func readOneAnimals(){

}

func updateAnimals(){

}

func deleteOneAnimals(){

}

func deleteManyAnimals(){

}