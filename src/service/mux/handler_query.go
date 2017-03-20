package mux

import (
	"encoding/json"
	"log"
	"net/http"
	"querybuilder"
	"service/db"
	"service/protocol"
)

func init() {
	handlerCollection["/nagp/records"] = handleQuery
}

func handleQuery(w http.ResponseWriter, r *http.Request, database db.Database) {

	qp := &protocol.QueryParam{}
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&qp)
	if err != nil {
		log.Printf("got error while decoding query param from body: %v\n", err)
		w.Write([]byte("internal error"))
		return
	}

	cities, err := database.QueryCities(qp.Query, querybuilder.Builder{})
	if err != nil {
		log.Printf("got error while finding cities Error : %v\n", err)
		w.Write([]byte("internal error"))
		return
	}

	data, err := json.Marshal(cities)
	if err != nil {
		log.Printf("got error while marshaling cities Error : %v\n", err)
		w.Write([]byte("internal error"))
		return
	}

	w.Write(data)

}
