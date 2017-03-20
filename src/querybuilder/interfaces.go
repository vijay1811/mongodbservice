package querybuilder

import (
	"gopkg.in/mgo.v2/bson"
)

type QueryBuilder interface {
	BuildQuery(logicalQuery string) (*bson.M, error)
}

type queryAble interface {
	getBSON() bson.M
}
