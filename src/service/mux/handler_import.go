package mux

import (
	"encoding/json"
	"log"
	"net/http"
	"service/db"
	"service/protocol"
)

func init() {
	handlerCollection["/nagp/import"] = handleImport
}

func handleImport(w http.ResponseWriter, r *http.Request, database db.Database) {

	ip := &protocol.ImportParam{}
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&ip)
	if err != nil {
		log.Printf("got error while decoding import param from body: %v\n", err)
		w.Write([]byte("internal error"))
		return
	}

	res, err := database.ImportCollection(ip)
	if err != nil {
		log.Printf("got error while importing info: %v, Error : %v\n", res, err)
		w.Write([]byte("internal error"))
		return
	}

	w.Write([]byte(res))

}
