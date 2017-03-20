package querybuilder

import (
	"gopkg.in/mgo.v2/bson"
)

type stackBson struct {
	data []*bson.M
}

func (s *stackBson) Push(val *bson.M) {
	s.data = append(s.data, val)
}

func (s *stackBson) Pop() *bson.M {
	if len(s.data) == 0 {
		return nil
	}
	retData := s.data[len(s.data)-1]
	s.data[len(s.data)-1] = nil
	s.data = s.data[:len(s.data)-1]

	return retData
}

func (s *stackBson) Top() *bson.M {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data)-1]
}

func (s *stackBson) Empty() bool {
	return len(s.data) == 0
}
