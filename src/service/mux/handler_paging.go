package mux

import (
	"encoding/json"
	"log"
	"net/http"
	"service/db"
	"service/protocol"
)

func init() {
	handlerCollection["/nagp/places"] = handlePaging
}

func handlePaging(w http.ResponseWriter, r *http.Request, database db.Database) {

	p := &protocol.PagingParam{}
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&p)
	if err != nil {
		log.Printf("got error while decoding paging param from body: %v\n", err)
		w.Write([]byte("internal error"))
		return
	}

	cities, err := database.CitiesPage(p)
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
