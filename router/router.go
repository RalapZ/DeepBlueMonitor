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

func MainFunc1(conf model.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		test, _ := ioutil.ReadAll(r.Body)
		var DataInfo []model.SkywalkInfo
		json.Unmarshal(test, &DataInfo)
		//fmt.Println(&conf)
		for _, Message := range DataInfo {
			common.SendMessage(Message, conf)
		}
	}
}

func MainFunc(conf *model.Config) {
	http.HandleFunc("/alarm", MainFunc1(*conf))
	http.Handle("/metric", promhttp.Handler())
	fmt.Printf("%#v", conf)
	http.ListenAndServe(fmt.Sprintf(":%s", conf.Listenport), nil)
}
