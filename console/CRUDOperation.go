package console

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
var db *sql.DB

func initDB() {
    var err error
    db, err = sql.Open("sqlite", "./database/animals.db")
    if err != nil {
        panic(err)
    }
}

func createAnimals() {
	if db == nil{
		initDB()
	}

	var createAnimal animal
	fmt.Print("Введите имя животного: ")
	fmt.Scan(&createAnimal.name)
	fmt.Print("Введите тип животного: ")
	fmt.Scan(&createAnimal.type_animals)
	fmt.Print("Введите породу животного: ")
	fmt.Scan(&createAnimal.breed)
	fmt.Print("Введите возраст животного: ")
	fmt.Scan(&createAnimal.age)
	fmt.Print("Введите пол животного: ")
	fmt.Scan(&createAnimal.gender)
	fmt.Print("Введите цвет животного: ")
	fmt.Scan(&createAnimal.color)
	result, err := db.Exec(`
		INSERT INTO animals (name, type_animals, breed, age, gender, color)
		VALUES (?, ?, ?, ?, ?, ?)`, createAnimal.name, createAnimal.type_animals, createAnimal.breed, createAnimal.age, createAnimal.gender, createAnimal.color)
	
	if err != nil{
		panic(err)
	}
	
	var createId, _ = result.LastInsertId()
	fmt.Println("Объект успешно вставлен с id - ", createId)
}

func readAllAnimals() {
	if db == nil {
		initDB()
	}

	rows, err := db.Query("SELECT * FROM animals")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

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
	if db == nil {
		initDB()
	}

	var id string
	fmt.Print("Укажите id: ")
	fmt.Scan(&id)

	rows, err := db.Query("SELECT * FROM animals WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

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
	readAllAnimals()
	var (
		id int
		lastAnimal animal
		updateAnimal animal
	)
	fmt.Print("Укажите id животного для удаления: ")
	fmt.Scan(&id)
	
	rows, err := db.Query("SELECT * FROM animals WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&lastAnimal.id, &lastAnimal.name, &lastAnimal.type_animals, &lastAnimal.breed, &lastAnimal.age, &lastAnimal.gender, &lastAnimal.color)
	} else {
		fmt.Println("Введенного id не существует в базе")
	}

	fmt.Printf("Укажите измененые атрибуты, если вы не хотите изменять определенный атрибут то укажите поле пустым\nВведите имя животного(прошлое значение: %s): ", lastAnimal.name)
	fmt.Scan(&updateAnimal.name)
	if updateAnimal.name == "" {updateAnimal.name = lastAnimal.name}
	fmt.Printf("Введите тип животного(прошлое значение: %s): ", lastAnimal.type_animals)
	fmt.Scan(&updateAnimal.type_animals)
	if updateAnimal.type_animals == "" {updateAnimal.type_animals = lastAnimal.type_animals}
	fmt.Printf("Введите породу животного(прошлое значение: %s): ", lastAnimal.breed)
	fmt.Scan(&updateAnimal.breed)
	if updateAnimal.breed == "" {updateAnimal.breed = lastAnimal.breed}
	fmt.Printf("Введите возраст животного(прошлое значение: %d): ", lastAnimal.age)
	fmt.Scan(&updateAnimal.age)
	if updateAnimal.age == 0 {updateAnimal.age = lastAnimal.age}
	fmt.Printf("Введите пол животного(прошлое значение: %s): ", lastAnimal.gender)
	fmt.Scan(&updateAnimal.gender)
	if updateAnimal.gender == "" {updateAnimal.gender = lastAnimal.gender}
	fmt.Printf("Введите цвет животного(прошлое значение: %s): ", lastAnimal.color)
	fmt.Scan(&updateAnimal.color)
	if updateAnimal.color == "" {updateAnimal.color = lastAnimal.color}

	result, err := db.Exec(`
		UPDATE animals SET name = ?, type_animals = ?, breed = ?,
		age = ?, gender = ?, color = ? WHERE id = ?`, 
		updateAnimal.name, updateAnimal.type_animals, updateAnimal.breed,
		updateAnimal.age, updateAnimal.gender, updateAnimal.color, id)
	
	if err != nil{
		panic(err)
	}
	var updateId, _ = result.RowsAffected()
	if(updateId == 0){
		fmt.Println("Ничего не изменилось")
	} else{
		fmt.Println("Изменения успешно сохранились")
	}
}

func deleteOneAnimals() {
	if db == nil {
		initDB()
	}

	var id string
	fmt.Print("Укажите id: ")
	fmt.Scan(&id)

	statement, err := db.Prepare("DELETE FROM animals WHERE id = ?")
	if err != nil {
		fmt.Println("Ошибка подготовки запроса:", err)
		return
	}

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
