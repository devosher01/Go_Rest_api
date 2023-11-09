package main

import (
	"net/http"

	"github.com/devosher01/go_gorm_api/db"
	"github.com/devosher01/go_gorm_api/models"
	"github.com/devosher01/go_gorm_api/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConection()
	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Task{})
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	// Tasks routes


	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.UpdateTaskHandler).Methods("PUT")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")
	http.ListenAndServe(":3000", r)
}
