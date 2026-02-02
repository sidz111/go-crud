package main

import (
	"go-mysql-crud/controller"
	"go-mysql-crud/db"
	"go-mysql-crud/repository"
	"log"
	"net/http"
)

func main() {
	database := db.ConnectDB()

	repo := &repository.EmployeeRepository{DB: database}
	ctrl := &controller.EmployeeController{Repo: repo}

	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			ctrl.CreateEmployee(w, r)
		case http.MethodGet:
			ctrl.GetEmployees(w, r)
		case http.MethodPut:
			ctrl.UpdateEmployee(w, r)
		case http.MethodDelete:
			ctrl.DeleteEmployee(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("🚀 Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
