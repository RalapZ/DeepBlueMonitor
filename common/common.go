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

func HttpRequest(HttpMethod string, Url string, data map[string]interface{}) ([]byte, error) {
	//func HttpRequest(data map[string]interface{},Url string) ([]byte,error){
	HeaderStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	req, err := http.NewRequest(HttpMethod, Url, bytes.NewReader(HeaderStr))
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func TokenGet(CorpId string, CorpSecret string) (string, error) {
	data := make(map[string]interface{})
	data["corpid"] = CorpId
	data["corpsecret"] = CorpSecret
	HttpMethod := "POST"
	Url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken"
	respbody, err := HttpRequest(HttpMethod, Url, data)
	if err != nil {
		return "", err
	}
	qistu := model.QiYchatStu{}
	err = json.Unmarshal(respbody, &qistu)
	if err != nil {
		return "", err
	}
	return qistu.Access_token, nil
}

func SendMessage(M *model.SkywalkInfo, conf *model.Config) {
	CorpId := conf.Tencent.Auth.CorpInfo
	CorpSecret := conf.Tencent.Auth.CorpSecret
	Token, err := TokenGet(CorpId, CorpSecret)
	if err != nil {
		panic(err.Error())
	}
	message := make(map[string]interface{})
	//fmt.Println(&conf)
	var user string
	for _, v := range conf.Tencent.User {
		user = v + "|" + user
	}
	message["touser"] = user
	message["msgtype"] = "text"
	message["agentid"] = conf.Tencent.Agentid
	str := "服务名称:" + M.Name + "\n告警类型:" + M.Scope + "\n告警规则:" + M.RuleName + "" + "\n告警信息:" + M.AlarmMessage
	message["text"] = map[string]interface{}{
		"content": str,
	}
	message["safe"] = "0"
	Url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + Token
	log.Println(str)
	HttpMethod := "POST"
	HttpBody, _ := HttpRequest(HttpMethod, Url, message)
	fmt.Println(string(HttpBody))
}
