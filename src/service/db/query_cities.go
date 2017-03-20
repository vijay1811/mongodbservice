package db

import (
	"fmt"
	"querybuilder"
)

func (Default) QueryCities(logicalQuery string, qb querybuilder.QueryBuilder) ([]*City, error) {

	bsonQuery, err := qb.BuildQuery(logicalQuery)
	if err != nil {
		return nil, err
	}

	query := session.DB("test").C("cities").Find(bsonQuery)
	var cities []*City
	n, err := query.Count()
	if err != nil {
		fmt.Printf("count: %v, error: %v\n", n, err)
	}

	err = query.All(&cities)
	if err != nil {
		return nil, err
	}

	return cities, err
}
