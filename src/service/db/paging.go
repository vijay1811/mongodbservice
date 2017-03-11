package db

import (
	"gopkg.in/mgo.v2/bson"
	"service/protocol"
)

func (Default) CitiesPage(p *protocol.PagingParam) ([]*City, error) {
	bsonQuery := bson.M{}

	query := session.DB("test").C("cities").Find(bsonQuery)

	skipLength := uint32(0)
	if p.PageIndex > 0 {
		skipLength = (p.PageIndex - 1) * p.PageSize
	}

	query = query.Skip(int(skipLength)).Limit(int(p.PageSize))
	var cities []*City
	err := query.All(&cities)
	if err != nil {
		return nil, err
	}

	return cities, err
}
