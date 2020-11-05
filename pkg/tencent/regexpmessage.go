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
//   Date  : 2020/10/27                                                                             //
//##################################################################################################//

package test

import (
	"fmt"
	"regexp"
)

func StrRegexp(SrcStr string, RegexpStr []string) string {
	for _, str := range RegexpStr {
		re := str + "[-,a-z,0-9]*"
		restr := regexp.MustCompile(re)
		FindStr := restr.FindAll([]byte(SrcStr), -1)
		if len(FindStr) != 0 {
			return string(FindStr[0])
		}
		//fmt.Println("1")
	}
	return ""
}

func main() {
	RegexpStr := []string{"crm-", "ssm-", "dlm-", "wx-"}
	SrcStr := "User in User to {GET}/interfaces/server/app/queryOrderInfo/app006/interAspect in wx-dmscloud-customer-service-prd1"
	str := StrRegexp(SrcStr, RegexpStr)
	fmt.Println(str)
}
