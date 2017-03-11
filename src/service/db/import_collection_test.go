package db

import (
	"service/protocol"
	"testing"
)

func TestCollectionImport(t *testing.T) {
	database, err := NewSession()
	if err != nil {
		t.Fatalf("error while starting the database, Err: %v\n", err)
	}

	defer func() {
		session.Close()
		session = nil
	}()

	cip := &protocol.ImportParam{
		Database:   "test",
		Collection: "restaurants",
		File:       "/home/nagarro/workspace/primer-dataset.json",
	}

	result, err := database.ImportCollection(cip)
	if err != nil {
		t.Fatalf("Error while importing: %+v,Error: \n", cip, err)
	}
	if result == "" {
		t.Fatalf("did not expect result to be empty \n")
	}
}
