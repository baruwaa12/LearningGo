package saddlepoints

import (
	"testing"
	"reflect"
)


func TestBuildMatrix(t *testing.T) {
	for _,test := range buildMatrixtestCases {
		var result = buildMatrix(test.row, test.column);
		if reflect.DeepEqual(result, test.matrix) == false {
				t.Fatalf("FAILED given column %d row %d result was %T", test.column, test.row, result)
		}
		t.Logf("PASSED %s", test.description)
	}
}
