package main

import (
	"net/http"

	"github.com/anibal-alpizar/go-gorm-fastapi/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":8080", r)
}
