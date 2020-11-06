package common

import (
	"fmt"
	"testing"
)

func TestExecute(t *testing.T) {
	s, e := TokenGet("ww97af1eab5d2add3c", "iq2IyxRcY3oCHHTFg2U2o3UQGHzXIWkKIAgfKQFdhxw")
	if e != nil {
		fmt.Errorf(e.Error())
	}
	fmt.Println(s)
}
