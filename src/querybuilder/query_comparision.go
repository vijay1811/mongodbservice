package querybuilder

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type comparisonQuery struct {
	op        string
	fieldName string
	value     interface{}
}

func newComparisonQuery(op, fieldName, val string) (*comparisonQuery, error) {
	op, ok := comOpConvMap[op]
	if !ok {
		return nil, fmt.Errorf("Operator not supported: %v\n", op)
	}
	value, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return nil, fmt.Errorf("Cannot convert to float: %v\n", err)
	}
	return &comparisonQuery{
		op:        op,
		fieldName: fieldName,
		value:     value,
	}, nil
}

func (q *comparisonQuery) getBSON() *bson.M {
	return &bson.M{
		q.fieldName: &bson.M{
			q.op: q.value,
		},
	}
}
