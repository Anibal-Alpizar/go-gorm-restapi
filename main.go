package main

import (
	"net/http"

	"github.com/anibal-alpizar/go-gorm-fastapi/db"
	"github.com/anibal-alpizar/go-gorm-fastapi/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":8080", r)
}
