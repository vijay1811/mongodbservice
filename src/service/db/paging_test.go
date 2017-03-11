package db

import (
	"service/protocol"
	"testing"
)

func TestCitiesPage(t *testing.T) {
	database, err := NewSession()
	if err != nil {
		t.Fatalf("error while starting the database, Err: %v\n", err)
	}

	defer func() {
		session.Close()
		session = nil
	}()

	p := &protocol.PagingParam{
		PageIndex: 1,
		PageSize:  10,
	}

	cities, err := database.CitiesPage(p)
	if err != nil {
		t.Fatalf("error while finding cities: %v\n", err)
	}

	if len(cities) != 10 {
		t.Fatalf("expected count: %v, actual count: %v\n", 9431, len(cities))
	}
}
