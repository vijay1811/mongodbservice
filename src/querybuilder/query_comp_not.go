package querybuilder

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type comparisonNotQuery struct {
	op1, op2  string
	fieldName string
	value     interface{}
}

func newComparisonNotQuery(op, fieldName, val string) (*comparisonNotQuery, error) {
	op1, ok := comOpConvMap[op[:0]]
	if !ok {
		return nil, fmt.Errorf("Operator not supported: %v\n", op)
	}
	op2, ok := comOpConvMap[op[1:]]
	if !ok {
		return nil, fmt.Errorf("Operator not supported: %v\n", op)
	}
	value, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return nil, fmt.Errorf("Cannot convert to float: %v\n", err)
	}
	return &comparisonNotQuery{
		op1:       op1,
		op2:       op2,
		fieldName: fieldName,
		value:     value,
	}, nil
}

func (q *comparisonNotQuery) getBSON() *bson.M {
	return &bson.M{
		q.fieldName: &bson.M{
			q.op1: &bson.M{
				q.op2: q.value,
			},
		},
	}
}
