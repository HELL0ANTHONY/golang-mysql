package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HELL0ANTHONY/crud-mysql-docker-go/controllers"
)

func run(p string) {
	error := http.ListenAndServe(p, nil)
	if error != nil {
		fmt.Printf("Error: %v", error)
		return
	}
	log.Println("Server running...")
}

func main() {
	const PORT = ":8080"
	http.HandleFunc("/", controllers.PrintTable)
	http.HandleFunc("/create", controllers.CreateUser)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.DeleteEmployee)
	http.HandleFunc("/edit", controllers.UpdateEmployee)

	run(PORT)
}
