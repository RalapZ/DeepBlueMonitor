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
	"context"
	"github.com/RalapZ/DeepBlueMonitor/common"
	"github.com/RalapZ/DeepBlueMonitor/model"
	"github.com/RalapZ/DeepBlueMonitor/router"
	"log"
	"os"
	"reflect"
	"strings"
)

var Conf model.Config

func InitEsClient() {
	var ctx = context.Background()
	//Url := []string{"http://10.10.8.151:9200/"}
	Url := []string{"http://10.10.8.151:9200/"}
	IndexName := common.IndexName
	Mapping := common.Mapping
	common.ESClientConn(Url)
	ESClient := common.ESClient
	exists, err := ESClient.IndexExists(IndexName).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		_, err := ESClient.CreateIndex(IndexName).BodyString(Mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
	}
}

func Execute() {
	//读取配置文件
	//filename := "C:/code/DeepBlueMonitor/conf/application.yaml"
	filename := "C:/code/DeepBlueMonitor/conf/application.yaml"
	var businessName []string
	Conf.ReadConfig(filename)
	//创建日志文件
	file, _ := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	//初始化es客户端
	InitEsClient()
	log.SetOutput(os.Stdout)
	//file,_:=os.OpenFile("catalina.log",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	//defer file.Close()
	//log.SetOutput(file)
	t := reflect.TypeOf(Conf.Tencent.BusinessType)
	for i := 0; i < t.NumField(); i++ {
		businessName = append(businessName, strings.ToLower(t.Field(i).Name))
	}
	router.MainFunc(&Conf, businessName)
}
