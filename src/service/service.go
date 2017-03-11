package main

import (
	"net/http"

	"service/db"
	"service/mux"
)

func main() {
	database, err := db.NewSession()
	if err != nil {
		panic(err)
	}

	handler := mux.NewHandler(database)

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	server.ListenAndServe()
}
