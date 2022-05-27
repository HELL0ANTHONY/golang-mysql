package controllers

import (
	"net/http"
	"text/template"

	"github.com/HELL0ANTHONY/crud-mysql-docker-go/connection"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))

func errorHandle(err error) {
	if err != nil {
		panic(err.Error())
	}
}

type Employee struct {
	Id    int
	Name  string
	Email string
}

func GetEmployees() []Employee {
	conn := connection.DBConnection()
	rows, err := conn.Query("SELECT * FROM employees")
	errorHandle(err)

	employee := Employee{}
	employees := []Employee{}

	for rows.Next() {
		var id int
		var name, email string
		err := rows.Scan(&id, &name, &email)
		errorHandle(err)
		employee.Id = id
		employee.Name = name
		employee.Email = email
		employees = append(employees, employee)
	}
	return employees
}

func PrintTable(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "table", GetEmployees()) // table = template name
	errorHandle(err)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "form", nil)
	errorHandle(err)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")

		ready := connection.DBConnection()
		insertValuesToDB, dbErr := ready.Prepare(
			`INSERT INTO employees (name, email) VALUES (?, ?);`,
		)
		errorHandle(dbErr)

		_, execErr := insertValuesToDB.Exec(name, email)
		errorHandle(execErr)
		http.Redirect(w, r, "/", 301)
	}
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	employeeId := r.URL.Query().Get("id")
	// log.Println("id", id)

	ready := connection.DBConnection()
	insertValuesToDB, dbErr := ready.Prepare(`DELETE * FROM employee WHERE id=?;`)
	errorHandle(dbErr)

	_, execErr := insertValuesToDB.Exec(employeeId)
	errorHandle(execErr)
	http.Redirect(w, r, "/", 301)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	employeeId := r.URL.Query().Get("id")
	conn := connection.DBConnection()
	row, err := conn.Query(`SELECT * FROM employees WHERE id=?;`, employeeId)
	errorHandle(err)
	employee := Employee{}
	for row.Next() {
		var id int
		var name, email string
		err := row.Scan(&id, &name, &email)
		errorHandle(err)
		employee.Id = id
		employee.Name = name
		employee.Email = email
	}
}
