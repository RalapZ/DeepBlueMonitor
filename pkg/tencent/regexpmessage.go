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

package main

import (
	"fmt"
	"regexp"
)

func StrRegexp(SrcStr string,RegexpStr []string) {
	for _,str:=range RegexpStr {
		re:=str+"[-,a-z,0-9]*"
		restr := regexp.MustCompile(re)
		FindStr :=restr.FindAll([]byte(SrcStr), -1)
		if len(FindStr)!=0{
			for k,i :=range(FindStr){
				fmt.Println(k,string(i),"\n")
			}
		}
		//if len(FindStr)!=0{
		//	return
		//}
	}
}

func main(){
	RegexpStr:=[]string{"crm","ssm","dlm","wx"}
	SrcStr:="User in User to {GET}/interfaces/server/app/queryOrderInfo/app006/interAspect in crm-dmscloud-customer-service-prd1"
	StrRegexp(SrcStr,RegexpStr)
	//fmt.Println(str)
}