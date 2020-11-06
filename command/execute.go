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

package command

import (
	"github.com/RalapZ/DeepBlueMonitor/model"
	"github.com/RalapZ/DeepBlueMonitor/router"
	"log"
	"os"
)

func Execute() {
	filename := "C:/code/DeepBlueMonitor/conf/application.yaml"
	var conf model.Config
	conf.ReadConfig(filename)
	file, _ := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	log.SetOutput(file)
	log.Fatal("asdfasdfasdf")
	//file,_:=os.OpenFile("catalina.log",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	//defer file.Close()
	//log.SetOutput(file)
	//log.Fatal("asdfasdfasdfasdfasdfasdfaf")
	router.MainFunc(&conf)
}
