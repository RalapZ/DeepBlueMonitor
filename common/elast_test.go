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
//   Date  : 2020/11/17                                                                             //
//##################################################################################################//
package common

import (
	"context"
	"fmt"
	"github.com/RalapZ/DeepBlueMonitor/model"
	"testing"
)

func TestElasticPost(t *testing.T) {
	var ctx = context.Background()
	Url := []string{"http://10.10.8.151:9200/"}
	ESClientConn(Url)
	fmt.Println("TEst", ESClient)
	info := model.SkywalkInfo{ScopeId: 222,
		Scope:         "myzonezhuRalap",
		Name:          "dlm-tasdfasdfest2",
		Id0:           "asdadf",
		Id1:           "asdfadfadsf",
		RuleName:      "service_resp_time_rule",
		AlarmMessage:  "1000ms in 3 minutes of last 10 minutes.",
		Starttimerecv: 159905164357,
		BInfo: model.BuinessInfo{BusinesType: "ssm",
			BusinesName: ""},
	}
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
	ElasticPostData(&info)
}
