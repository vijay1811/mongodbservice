package db

import (
	"testing"
)

func TestNewSession(t *testing.T) {
	database, err := NewSession()
	if err != nil {
		t.Fatalf("error while starting the database, Err: %v\n", err)
	}
	defer func() {
		session.Close()
		session = nil
	}()

	if database == nil {
		t.Fatalf("did not expect database to be nil %v\n", err)
	}

	if session == nil {
		t.Fatalf("expected session to be initialize\n")
	}

	if err := session.Ping(); err != nil {
		t.Fatalf("got error while pinging to MongoDB server Err: %v\n", err)
	}

}
