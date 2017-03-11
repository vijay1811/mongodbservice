package mux

import (
	"net/http"
	"service/db"
)

var handlerCollection = make(map[string]func(http.ResponseWriter, *http.Request, db.Database))
