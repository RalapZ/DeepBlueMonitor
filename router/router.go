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
	"fmt"
	"github.com/RalapZ/DeepBlueMonitor/common"
	"github.com/RalapZ/DeepBlueMonitor/model"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"net/http"
)

func QiyeChatMessage(conf model.Config, businessName []string) func(w http.ResponseWriter, r *http.Request) {
	//temptest.Myzone()
	return func(w http.ResponseWriter, r *http.Request) {
		test, _ := ioutil.ReadAll(r.Body)
		var DataInfo []model.SkywalkInfo
		json.Unmarshal(test, &DataInfo)
		fmt.Printf("qiyechatmessage   %#v", DataInfo)
		for _, Message := range DataInfo {
			common.SendMessage(&Message, &conf, businessName)
		}
	}
}

func MainFunc(conf *model.Config, businessName []string) {
	http.HandleFunc("/alarm", QiyeChatMessage(*conf, businessName))
	http.Handle("/metric", promhttp.Handler())
	//fmt.Printf("%#v", conf)
	http.ListenAndServe(fmt.Sprintf(":%s", conf.Listenport), nil)
}
