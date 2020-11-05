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

package main

import (
	"fmt"
	"github.com/RalapZ/DeepBlueMonitor/model"
	"github.com/RalapZ/DeepBlueMonitor/router"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	//common.Execute()
	filename := "C:/code/DeepBlueMonitor/conf/application.yaml"
	var conf model.Config
	//str,_:=os.Getwd()
	//fmt.Println(string(str))
	conf.ReadConfig(filename)
	//fmt.Println(conf)
	//fmt.Println(conf.Listenport)
	http.HandleFunc("/alarm", router.MainFunc(conf))
	http.Handle("/metric", promhttp.Handler())
	fmt.Println(conf)
	http.ListenAndServe(fmt.Sprintf(":%s", conf.Listenport), nil)
}
