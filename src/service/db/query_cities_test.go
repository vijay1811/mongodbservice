package db

import (
	"querybuilder"
	"testing"
)

func TestQueryCities(t *testing.T) {
	database, err := NewSession()
	if err != nil {
		t.Fatalf("error while starting the database, Err: %v\n", err)
	}

	defer func() {
		session.Close()
		session = nil
	}()

	cities, err := database.QueryCities(" pop > 27000 ", querybuilder.Builder{})
	if err != nil {
		t.Fatalf("error while finding cities: %v\n", err)
	}

	if len(cities) != 2765 {
		t.Fatalf("expected count: %v, actual count: %v\n", 2765, len(cities))
	}
}
