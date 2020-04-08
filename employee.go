/**
 * Original from
 * https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html
 */
package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var tmpl = template.Must(template.ParseGlob("views/*"))

/**
 * Employee struct
 */
type Employee struct {
	Id   int
	Name string
	City string
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", index)
	http.HandleFunc("/show", show)
	http.HandleFunc("/new", new)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", update)
	http.ListenAndServe(":8080", nil)
}

func dbConn() (db *sql.DB) {
	var dbDriver string = "mysql"
	var dbUser string = "root"
	var dbPass string = "123456"
	var dbName string = "go_employee"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}

/**
 * Index List employee
 */
func index(w http.ResponseWriter, r *http.Request) {
	var db *sql.DB = dbConn()
	var sql string = "SELECT * FROM employee ORDER BY id DESC"
	selDb, err := db.Query(sql)

	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}
	res := []Employee{}

	for selDb.Next() {
		var id int
		var name, city string
		err = selDb.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}

		emp.Id = id
		emp.Name = name
		emp.City = city

		res = append(res, emp)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	// fmt.Println(res)
	defer db.Close()
}

/**
 * Show Employee detail
 */
func show(w http.ResponseWriter, r *http.Request) {
	var db *sql.DB = dbConn()

	var nId string = r.URL.Query().Get("id")
	selDb, err := db.Query("SELECT * FROM employee WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	var emp Employee = Employee{}
	for selDb.Next() {
		var id int
		var name, city string
		err = selDb.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}

		emp.Id = id
		emp.Name = name
		emp.City = city
	}

	tmpl.ExecuteTemplate(w, "Show", emp)

	defer db.Close()
}

func new(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		insForm, err := db.Prepare("INSERT INTO employee(name, city) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city)
		log.Println("INSERT: Name: " + name + " | City: " + city)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE employee SET name=?, city=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city, id)
		log.Println("UPDATE: Name: " + name + " | City: " + city)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM employee WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM employee WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
	}
	tmpl.ExecuteTemplate(w, "Edit", emp)
	defer db.Close()
}
