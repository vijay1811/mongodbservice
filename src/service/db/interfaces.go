package db

import (
	"service/protocol"
)

type Database interface {
	// ImportCollection imports given data in a given database and collection
	// return string and an error
	ImportCollection(cip *protocol.ImportParam) (string, error)

	// ImportCollection takes CitiesParam returns cities matching that parameters
	FindCities(c *protocol.CitiesParam) ([]*City, error)

	// ImportCollection takes PagingParam returns cities in the given page number and index
	CitiesPage(p *protocol.PagingParam) ([]*City, error)
}
