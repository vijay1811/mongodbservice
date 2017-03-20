package querybuilder

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type Builder struct{}

func (Builder) BuildQuery(logicalQuery string) (*bson.M, error) {
	return buildQuery(logicalQuery)
}

func buildQuery(logicalQuery string) (*bson.M, error) {
	postfix, err := infixToPostfix(logicalQuery)
	if err != nil {
		return nil, err
	}
	fields := strings.Split(postfix, " ")
	strStack := &stack{}
	bsStack := &stackBson{}
	for _, field := range fields {
		if isOperator(field) {
			switch operatorType(field) {
			case opType_comparison:
				if strStack.Empty() {
					return nil, fmt.Errorf("Did not expect stack to be empty")
				}
				docVal := strStack.Pop()
				if strStack.Empty() {
					return nil, fmt.Errorf("Did not expect stack to be empty")
				}
				docField := strStack.Pop()
				q, err := newComparisonQuery(field, docField, docVal)
				if err != nil {
					return nil, err
				}
				bsStack.Push(q.getBSON())
			case opType_NotComp:
				if strStack.Empty() {
					return nil, fmt.Errorf("Did not expect stack to be empty")
				}
				docVal := strStack.Pop()
				if strStack.Empty() {
					return nil, fmt.Errorf("Did not expect stack to be empty")
				}
				docField := strStack.Pop()
				q, err := newComparisonNotQuery(field, docField, docVal)
				if err != nil {
					return nil, err
				}
				bsStack.Push(q.getBSON())
			case opType_logical:
				if bsStack.Empty() {
					return nil, fmt.Errorf("Did not expect stack to be empty")
				}
				bson1 := bsStack.Pop()
				if bsStack.Empty() {
					return nil, fmt.Errorf("Did not expect stack to be empty")
				}
				bson2 := bsStack.Pop()
				q, err := newLogicalQuery(field, bson1, bson2)
				if err != nil {
					return nil, err
				}
				bsStack.Push(q.getBSON())
			}
		} else {
			strStack.Push(field)
		}
	}
	return bsStack.Pop(), nil
}
