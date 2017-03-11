package mux

import (
	"encoding/json"
	"log"
	"net/http"
	"service/db"
	"service/protocol"
)

func init() {
	handlerCollection["/nagp/cities"] = handleCities
}

func handleCities(w http.ResponseWriter, r *http.Request, database db.Database) {

	cp := &protocol.CitiesParam{}
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&cp)
	if err != nil {
		log.Printf("got error while decoding cities param from body: %v\n", err)
		w.Write([]byte("internal error"))
		return
	}

	cities, err := database.FindCities(cp)
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
