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
//   Date  : 2020/11/05                                                                             //
//##################################################################################################//
package router

import (
	"encoding/json"
	"github.com/RalapZ/DeepBlueMonitor/common"
	"github.com/RalapZ/DeepBlueMonitor/model"
	"io/ioutil"
	"net/http"
)

func MainFunc(conf model.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		test, _ := ioutil.ReadAll(r.Body)
		var DataInfo []model.SkywalkInfo
		json.Unmarshal(test, &DataInfo)
		//fmt.Println(&conf)
		for _, Message := range DataInfo {
			//corpid := "ww97af1eab5d2add3c"
			//corpsecret := "iq2IyxRcY3oCHHTFg2U2o3UQGHzXIWkKIAgfKQFdhxw"
			//SendMessage(corpid, corpsecret, Message,conf)
			common.SendMessage(Message, conf)
		}
	}
}
