package db

import (
	mgo "gopkg.in/mgo.v2"
)

const (
	//TODO make this configurable
	mongoDBServerUrl = "localhost"
)

var session *mgo.Session

type Default struct{}

// NewSession will returns an implementation of Database interface and an error
// if we fail to make session with database
func NewSession() (*Default, error) {

	// For now we want database to be singleton so we only making a session if session is
	// not initialized.
	if session == nil {
		s, err := mgo.Dial(mongoDBServerUrl)
		if err != nil {
			return nil, err
		}
		// we have made a session successfully now initialize our session variable
		session = s
	}

	return &Default{}, nil
}
