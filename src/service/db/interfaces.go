package db

import (
	"querybuilder"
	"service/protocol"
)

type Database interface {
	// ImportCollection imports given data in a given database and collection
	// return string and an error
	ImportCollection(cip *protocol.ImportParam) (string, error)

	// FindCities takes CitiesParam returns cities matching that parameters
	FindCities(c *protocol.CitiesParam) ([]*City, error)

	// CitiesPage takes PagingParam returns cities in the given page number and index
	CitiesPage(p *protocol.PagingParam) ([]*City, error)

	// QueryCities takes logical query string and returns cities matching the query
	QueryCities(logicalQuery string, qb querybuilder.QueryBuilder) ([]*City, error)
}
