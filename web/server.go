package web

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "modernc.org/sqlite"
)
type Animal struct {
    Id int
    Name string
    TypeAnimals string
    Breed string
    Age int
    Gender string
    Color string
}
var db *sql.DB

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    fmt.Printf("Id объекта для удаления: %s", id)
    _, err := db.Exec("DELETE FROM animals WHERE id = ?", id)
    if err != nil{
        log.Println(err)
    }
     
    http.Redirect(w, r, "/", 301)
}

func EditPage(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
 
    row := db.QueryRow("SELECT * FROM animals WHERE id = ?", id)
    anim := Animal{}
    err := row.Scan(&anim.Id, &anim.Name, &anim.TypeAnimals, &anim.Breed, &anim.Age, &anim.Gender, &anim.Color)
    if err != nil{
        log.Println(err)
        http.Error(w, http.StatusText(404), http.StatusNotFound)
    }else{
        tmpl, _ := template.ParseFiles("templates/edit.html")
        tmpl.Execute(w, anim)
    }
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        log.Println(err)
    }
    id := r.FormValue("id")
    name := r.FormValue("name")
    type_animals := r.FormValue("type_animals")
    breed := r.FormValue("breed")
    age := r.FormValue("age")
    gender := r.FormValue("gender")
    color := r.FormValue("color")
 
    _, err = db.Exec(`
		UPDATE animals SET name = ?, type_animals = ?, breed = ?,
		age = ?, gender = ?, color = ? WHERE id = ?`, 
		name, type_animals, breed,
		age, gender, color, id)
 
    if err != nil {
        log.Println(err)
    }
    http.Redirect(w, r, "/", 301)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
 
        err := r.ParseForm()
        if err != nil {
            log.Println(err)
        }
        name := r.FormValue("name")
        type_animals := r.FormValue("type_animals")
        breed := r.FormValue("breed")
        age := r.FormValue("age")
        gender := r.FormValue("gender")
        color := r.FormValue("color")
 
        _, err = db.Exec("INSERT INTO animals (name, type_animals, breed, age, gender, color) values (?, ?, ?, ?, ?, ?)", 
          name, type_animals, breed, age, gender, color)
 
        if err != nil {
            log.Println(err)
          }
        http.Redirect(w, r, "/", 301)
    }else{
        http.ServeFile(w,r, "templates/create.html")
    }
    // После успешного создания животного
    fmt.Fprintf(w, `
    <script>
        window.onload = function() {
            setTimeout(() => window.close(), 300);
        }
    </script>
    <div style="text-align: center; padding: 50px;">
        <h3>Животное успешно добавлено!</h3>
        <p>Окно закроется автоматически...</p>
        <button onclick="window.close()">Закрыть сейчас</button>
    </div>
    `)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT * FROM animals")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    animals := []Animal{}

    for rows.Next() {
        a := Animal{}
        err := rows.Scan(&a.Id, &a.Name, &a.TypeAnimals, &a.Breed, &a.Age, &a.Gender, &a.Color)
        if err != nil {
            log.Println(err)
            continue
        }
        animals = append(animals, a)
    }

    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    } 

    err = tmpl.Execute(w, animals)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
 
func RunServer() {
    os.Chdir("web")
    database, err := sql.Open("sqlite", "../database/animals.db")
    if err != nil {
        log.Fatal(err)
    }
    db = database
    defer database.Close()

    // Проверяем соединение с базой
    err = db.Ping()
    if err != nil {
        log.Fatal("Database connection failed:", err)
    }
    log.Println("Database connected successfully")

    router := mux.NewRouter()
    router.HandleFunc("/", IndexHandler)
    router.HandleFunc("/create", CreateHandler)
    router.HandleFunc("/edit/{id:[0-9]+}", EditPage).Methods("GET")
    router.HandleFunc("/edit/{id:[0-9]+}", EditHandler).Methods("POST")
    router.HandleFunc("/delete/{id:[0-9]+}", DeleteHandler)

    http.Handle("/",router)

    fmt.Println("Server is listening on :8181...")
    log.Fatal(http.ListenAndServe(":8181", nil))
}