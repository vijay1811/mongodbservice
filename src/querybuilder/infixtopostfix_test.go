package querybuilder

import (
	"testing"
)

func TestInfixToPostfix(t *testing.T) {
	postfix, err := infixToPostfix(" ( ( x > 5 ) && ( y < 10 || y > 30 ) ) || ( x < 30 ) ")
	if err != nil {
		t.Fatalf("did not expect error here: %v\n", err)
	}
	expPostfix := "x 5 > y 10 < y 30 > || && x 30 < ||"
	if expPostfix != postfix {
		t.Fatalf("expected: %s, act: %s", expPostfix, postfix)
	}
}
