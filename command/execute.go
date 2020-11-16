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
	"fmt"
	"github.com/RalapZ/DeepBlueMonitor/model"
	"github.com/RalapZ/DeepBlueMonitor/router"
	"log"
	"os"
	"reflect"
	"strings"
)

func Execute() {
	filename := "C:/code/DeepBlueMonitor/conf/application.yaml"
	var conf model.Config
	var businessName []string
	conf.ReadConfig(filename)
	//fmt.Println(conf)
	file, _ := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	log.SetOutput(os.Stdout)
	log.Println("asdfasdfasdf")
	//file,_:=os.OpenFile("catalina.log",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	//defer file.Close()
	//log.SetOutput(file)
	t := reflect.TypeOf(conf.Tencent.BusinessType)
	for i := 0; i < t.NumField(); i++ {
		businessName = append(businessName, strings.ToLower(t.Field(i).Name))
	}
	fmt.Println(businessName)

	router.MainFunc(&conf, businessName)
}
