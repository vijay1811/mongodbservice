package querybuilder

import (
	"fmt"
	"strings"
)

func infixToPostfix(infix string) (string, error) {
	infix = strings.TrimSpace(infix)
	fields := strings.Split(infix, " ")
	var postfix []string
	st := &stack{}
	for _, field := range fields {
		if isOperator(field) {
			if field == operator_CloseBracket {
				op := st.Pop()
				for op != operator_startBracket {
					if st.Empty() {
						return "", fmt.Errorf("invalid query unbalanced brackets")
					}
					postfix = append(postfix, op)
					op = st.Pop()
				}
			} else if field == operator_startBracket {
				st.Push(field)
			} else {
				currentOpPrirority := priority(field)
				op := st.Top()
				for (priority(op) <= currentOpPrirority) && !st.Empty() {
					postfix = append(postfix, st.Pop())
					op = st.Top()
				}
				st.Push(field)
			}
		} else {
			postfix = append(postfix, field)
		}
	}
	for !st.Empty() {
		postfix = append(postfix, st.Pop())
	}
	return strings.Join(postfix, " "), nil
}
