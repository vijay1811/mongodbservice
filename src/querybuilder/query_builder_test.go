package querybuilder

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestBuildQuery(t *testing.T) {
	bsonQuery, err := buildQuery(" ( ( x > 5.6777 ) && ( y < 10 || y > 30 ) ) || ( x < 30 ) ")
	if err != nil {
		t.Fatalf("did not expect error here: %v\n", err)
	}
	fmt.Println(bsonQuery)
	json, _ := bson.MarshalJSON(bsonQuery)
	fmt.Println(string(json))

}
