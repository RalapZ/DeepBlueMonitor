//##################################################################################################//
//                   			         ┌─┐       ┌─┐ + +                                          //
//                   			      ┌──┘ ┴───────┘ ┴──┐++                                         //
//                   			      │       ───       │++ + + +                                   //
//                   			      ███████───███████ │+                                          //
//                   			      │       ─┴─       │+                                          //
//                   			      └───┐         ┌───┘                                           //
//                   			          │         │   + +                                         //
//                   			          │         └──────────────┐                                //
//                   			          │                        ├─┐                              //
//                   			          │                        ┌─┘                              //
//                   			          └─┐  ┐  ┌───────┬──┐  ┌──┘  + + + +                       //
//                   			            │ ─┤ ─┤       │ ─┤ ─┤                                   //
//                   			            └──┴──┘       └──┴──┘  + + + +                          //
//                   			      神兽出没               永无BUG                                 //
//   Author: Ralap                                                                                  //
//   Date  : 2020/11/13                                                                             //
//##################################################################################################//
package common

import (
	"errors"
	"regexp"
)

func StrRegexp(SrcStr string, RegexpStr []string) (string, string, error) {
	for _, buinesType := range RegexpStr {
		re := "(^|\\s)" + buinesType + "[-,a-z,0-9]*"
		restr := regexp.MustCompile(re)
		FindStr := restr.FindAll([]byte(SrcStr), 1)
		if len(FindStr) != 0 {
			return buinesType, string(FindStr[0]), nil
		}
	}
	return "", "", errors.New("not found")
}
