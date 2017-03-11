package mux

import (
	"fmt"
	"net/http"
	"service/db"
)

type Handler struct {
	database db.Database
}

func NewHandler(database db.Database) *Handler {
	return &Handler{
		database: database,
	}
}

//comment
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.String())
	if handle, ok := handlerCollection[r.URL.String()]; ok {
		handle(w, r, h.database)
		return
	}
	//TODO : Figure out what to do if url is not found
}
