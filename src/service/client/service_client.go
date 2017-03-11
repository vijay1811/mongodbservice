package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"service/protocol"
)

func main() {
	/*ip := protocol.ImportParam{
		Database:   "test",
		Collection: "cities",
		File:       `/home/nagarro/Downloads/Mongo DB Assignment/ctt data set/zips.json`,
	}*/

	/*cp := protocol.CitiesParam{
		Longitude: -74.0,
		Latitude:  40.74,
		Radius:    10.0,
	}*/

	p := &protocol.PagingParam{
		PageIndex: 10,
		PageSize:  10,
	}

	data, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:8080/nagp/places", "application/json", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

}
