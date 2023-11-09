package routes

import (
	"encoding/json"
	"net/http"

	"github.com/devosher01/go_gorm_api/db"
	"github.com/devosher01/go_gorm_api/models"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	muxVars := mux.Vars(r)

	db.DB.First(&task, muxVars["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("Task not found"))
		return
	}
	json.NewEncoder(w).Encode(task)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	// Decodifica el cuerpo de la solicitud en la tarea
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding request body: " + err.Error()))
		return
	}

	// Crea la tarea
	createdTask := db.DB.Create(&task)
	err = createdTask.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte("Error creating task: " + err.Error()))
		return
	}

	json.NewEncoder(w).Encode(task)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	muxVars := mux.Vars(r)

	db.DB.First(&task, muxVars["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("Task not found"))
		return
	}

	var newTask models.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding request body: " + err.Error()))
		return
	}

	db.DB.Model(&task).Updates(newTask)
	json.NewEncoder(w).Encode(task)

}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	muxVars := mux.Vars(r)

	db.DB.First(&task, muxVars["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("Task not found"))
		return
	}

	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusOK) // 200
	w.Write([]byte("Task deleted"))

}
