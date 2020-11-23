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
//   Date  : 2020/11/13                                                                             //
//##################################################################################################//
package common

import (
	"context"
	"fmt"
	"github.com/RalapZ/DeepBlueMonitor/model"
	"github.com/olivere/elastic/v7"
	"time"
)

var indexName = "skywalkinginfor"

const mapping = `
{
    "mappings": {
        "properties": {
            "StartTime": {
                "type": "date"
            },
            "ScopeId": {
                "type": "text"
            },
			"Scope": {
                "type": "text"
            },
			"BInfo":{
				"properties": {
					"BusinesType": {
						"type": "text"
					},
					"BusinesName": {
						"type": "text"
					}
				}
			},
			"Name": {
                "type": "text"
            },
			"id0": {
                "type": "text"
            },
			"id1": {
                "type": "text"
            },
			"RuleName": {
                "type": "text"
            },
            "AlarmMessage": {
                "type": "text"
            }
        }
    }
}`

func ElasticPost(info *model.SkywalkInfo, esclient *elastic.Client) {
	var ctx = context.Background()
	//Url := []string{"http://10.10.8.151:9200/"}
	//esclient, err := elastic.NewClient(elastic.SetURL(Url...), elastic.SetSniff(false))
	//if err != nil {
	//	panic(err)
	//}
	//info := model.SkywalkInfo{ScopeId: 222,
	//	Scope:        "asdfasdf232ewe",
	//	Name:         "ssm-test2",
	//	Id0:          "",
	//	Id1:          "",
	//	RuleName:     "service_resp_time_rule",
	//	AlarmMessage: "1000ms in 3 minutes of last 10 minutes.",
	//	BInfo: model.BuinessInfo{BusinesType: "ssm",
	//		BusinesName: ""},
	//}
	var t int64 = 1606113400
	info.StartTime = time.Unix(t, 0)
	//exists, err := esclient.IndexExists(indexName).Do(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//if !exists {
	//	_, err := esclient.CreateIndex(indexName).BodyString(mapping).Do(ctx)
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	tn := time.Now()
	info1, err := esclient.Index().
		Index(indexName).
		BodyJson(info).
		//Refresh("wait_for").
		Do(ctx)
	tn1 := time.Since(tn)
	fmt.Println(tn1)
	if err != nil {
		fmt.Println(info1)
		panic(err)
	}
}
