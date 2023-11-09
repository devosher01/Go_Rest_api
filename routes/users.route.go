package routes

import (
	"encoding/json"
	"net/http"

	"github.com/devosher01/go_gorm_api/db"
	"github.com/devosher01/go_gorm_api/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound) //404
		w.Write([]byte("User Not Found"))
		return
	}
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	json.NewEncoder(w).Encode(&user)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte("Error Creating User" + err.Error()))
		// json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(&user)

	//w.Write([]byte("Create User"))

}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound) //404
		w.Write([]byte("User Not Found"))
		return
	}

	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte("Error Creating User" + err.Error()))
		// json.NewEncoder(w).Encode(err)
		return
	}

	
	// Actualiza el usuario
	db.DB.Model(&user).Updates(newUser)

	json.NewEncoder(w).Encode(&user)

}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	// Busca el usuario por ID
	db.DB.First(&user, params["id"])

	// Si el usuario no existe, devuelve un error
	if user.ID == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Elimina el usuario
	result := db.DB.Unscoped().Delete(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return

	}
	w.WriteHeader(http.StatusOK) //20

}
