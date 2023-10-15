package main

import (
	"net/http"

	"github.com/gerachinchillar/gorm-api/db"
	"github.com/gerachinchillar/gorm-api/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)
}
