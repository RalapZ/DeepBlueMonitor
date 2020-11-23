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
	"github.com/RalapZ/DeepBlueMonitor/model"
	"github.com/olivere/elastic/v7"
	"testing"
)

func TestElasticPost(t *testing.T) {
	var ctx = context.Background()
	Url := []string{"http://10.10.8.151:9200/"}
	esclient, err := elastic.NewClient(elastic.SetURL(Url...), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	info := model.SkywalkInfo{ScopeId: 222,
		Scope:        "madfasdf222212222222",
		Name:         "crm-test2",
		Id0:          "asdadf",
		Id1:          "asdfadfadsf",
		RuleName:     "service_resp_time_rule",
		AlarmMessage: "1000ms in 3 minutes of last 10 minutes.",
		BInfo: model.BuinessInfo{BusinesType: "ssm",
			BusinesName: ""},
	}
	exists, err := esclient.IndexExists(indexName).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		_, err := esclient.CreateIndex(indexName).BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
	}
	ElasticPost(&info, esclient)
}
