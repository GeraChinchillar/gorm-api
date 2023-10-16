package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gerachinchillar/gorm-api/db"
	"github.com/gerachinchillar/gorm-api/model"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}
