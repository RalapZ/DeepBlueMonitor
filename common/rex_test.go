package common

import (
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	//B:=model.BuinessStruct{[]string{"ZHURONGJIA"},[]string{"ZHURONGJIA"},[]string{"ZHURONGJIA"},[]string{"ZHURONGJIA"}}
	srcstr := []string{"crm", "dlm", "ssm"}
	var buinessnamestring = "crmsdfasdfasdf"
	v, b, err := StrRegexp(buinessnamestring, srcstr)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(strings.ToUpper(v), b)

}
