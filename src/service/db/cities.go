package db

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"service/protocol"
)

type City struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	City  string        `bson:"city"`
	Loc   []float64     `bson:"loc"`
	Pop   uint32        `bson:"pop"`
	State string        `bson:"state"`
}

func (Default) FindCities(c *protocol.CitiesParam) ([]*City, error) {

	bsonQuery := bson.M{"loc": bson.M{
		"$geoWithin": bson.M{
			"$center": []interface{}{
				[]float64{c.Longitude, c.Latitude},
				c.Radius,
			},
		}}}

	query := session.DB("test").C("cities").Find(bsonQuery)
	var cities []*City
	n, err := query.Count()
	if err != nil {
		fmt.Printf("count: %v, error: %v\n", n, err)
	}
	if c.Sort != "" {
		query = query.Sort(c.Sort)
	}
	if c.Limit != 0 {
		query = query.Limit(int(c.Limit))
	}

	err = query.All(&cities)
	if err != nil {
		return nil, err
	}

	return cities, err
}
