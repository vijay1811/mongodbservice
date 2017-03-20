package querybuilder

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type logicalQuery struct {
	op               string
	subDoc1, subDoc2 *bson.M
}

func newLogicalQuery(op string, subDoc1, subDoc2 *bson.M) (*logicalQuery, error) {
	op, ok := comOpConvMap[op]
	if !ok {
		return nil, fmt.Errorf("Operator not supported: %v\n", op)
	}

	return &logicalQuery{
		op:      op,
		subDoc1: subDoc1,
		subDoc2: subDoc2,
	}, nil
}

func (q *logicalQuery) getBSON() *bson.M {
	return &bson.M{
		q.op: []*bson.M{
			q.subDoc1,
			q.subDoc2,
		},
	}
}
