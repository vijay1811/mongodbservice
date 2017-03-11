package db

import (
	"service/protocol"
	"testing"
)

func TestFindCities(t *testing.T) {
	database, err := NewSession()
	if err != nil {
		t.Fatalf("error while starting the database, Err: %v\n", err)
	}

	defer func() {
		session.Close()
		session = nil
	}()
	c := &protocol.CitiesParam{
		Longitude: -74.0,
		Latitude:  40.74,
		Radius:    10.0,
	}

	cities, err := database.FindCities(c)
	if err != nil {
		t.Fatalf("error while finding cities: %v\n", err)
	}

	if len(cities) != 9431 {
		t.Fatalf("expected count: %v, actual count: %v\n", 9431, len(cities))
	}
}
