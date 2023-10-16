package main

import (
	"net/http"

	"github.com/gerachinchillar/gorm-api/db"
	"github.com/gerachinchillar/gorm-api/model"
	"github.com/gerachinchillar/gorm-api/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(model.Task{})
	db.DB.AutoMigrate(model.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
