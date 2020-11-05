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

package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/RalapZ/DeepBlueMonitor/model"
	"io/ioutil"
	"log"
	"net/http"
)

func TokenGet(CorpId string, CorpSecret string) string {
	client := &http.Client{}
	data := make(map[string]interface{})
	data["corpid"] = CorpId
	data["corpsecret"] = CorpSecret
	HeaderStr, err := json.Marshal(data)
	if err != nil {
		return "nil"
	}
	req, _ := http.NewRequest("POST", "https://qyapi.weixin.qq.com/cgi-bin/gettoken", bytes.NewReader(HeaderStr))
	defer req.Body.Close()
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	qistu := model.QiYchatStu{}
	_ = json.Unmarshal(body, &qistu)
	return qistu.Access_token
}

func SendMessage(M model.SkywalkInfo, conf model.Config) {
	CorpId := conf.Tencent.Auth.CorpInfo
	CorpSecret := conf.Tencent.Auth.CorpSecret
	Token := TokenGet(CorpId, CorpSecret)
	client := &http.Client{}
	message := make(map[string]interface{})
	//fmt.Println(&conf)
	var user string
	for _, v := range conf.Tencent.User {
		user = v + "|" + user
	}
	//fmt.Println(user)
	message["touser"] = user
	message["msgtype"] = "text"
	message["agentid"] = conf.Tencent.Agentid
	str := "服务名称:" + M.Name + "\n告警类型:" + M.Scope + "\n告警规则:" + M.RuleName + "" + "\n告警信息:" + M.AlarmMessage
	message["text"] = map[string]interface{}{
		"content": str,
	}
	log.Println(str)
	message["safe"] = "0"
	Mdata, _ := json.Marshal(message)
	urlR := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + Token
	req1, _ := http.NewRequest("POST", urlR, bytes.NewReader(Mdata))
	defer req1.Body.Close()
	//fmt.Println(req1.Header)
	resp1, _ := client.Do(req1)
	MMessage, _ := ioutil.ReadAll(resp1.Body)
	//fmt.Println("=============")
	fmt.Println(string(MMessage))
}
