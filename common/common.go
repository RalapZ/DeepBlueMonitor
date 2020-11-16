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
	fmt.Println(data)
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
func BusinessUser(buiness string, conf *model.Config) []string {
	var user []string
	switch buiness {
	case "crm", "CRM":
		user = conf.Tencent.BusinessType.CRM
	case "ssm", "SSM":
		user = conf.Tencent.BusinessType.SSM
	case "dlm", "DLM":
		user = conf.Tencent.BusinessType.SSM
	case "VX", "vx":
		user = conf.Tencent.BusinessType.VX
	}
	return user
}

func SendMessage(M *model.SkywalkInfo, conf *model.Config, businessNameSlice []string) {
	CorpId := conf.Tencent.Auth.CorpInfo
	CorpSecret := conf.Tencent.Auth.CorpSecret
	Token, err := TokenGet(CorpId, CorpSecret)
	//bu := []string{"crm", "dlm", "ssm","vx"}
	fmt.Println("SendMessage", businessNameSlice)
	fmt.Println(M.Name)
	businessType, businessName, err := StrRegexp(M.Name, businessNameSlice)
	fmt.Println("StrRegexp", businessType, businessName)
	//businessType=strings.ToUpper(businessType)
	//usertemp:=[]string{}
	//switch businessType{
	//case "crm","CRM":
	//	usertemp=conf.Tencent.BusinessType.CRM
	//case "dlm","DLM":
	//	usertemp=conf.Tencent.BusinessType.DLM
	//case "ssm","SSM":
	//	usertemp=conf.Tencent.BusinessType.SSM
	//case "vx","VX":
	//	usertemp=conf.Tencent.BusinessType.VX
	//}
	//fmt.Println(usertemp)
	//fmt.Println(conf.Tencent.BusinessType.)
	if err != nil {
		panic(err.Error())
	}
	data := make(map[string]interface{})
	var buinessuser string
	usertemp := BusinessUser(businessType, conf)
	fmt.Println(usertemp, businessName)
	for _, v := range usertemp {
		buinessuser = v + "|" + buinessuser
	}
	//usertemp:=
	data["businessName"] = businessName
	data["touser"] = buinessuser
	data["msgtype"] = "text"
	data["agentid"] = conf.Tencent.Agentid
	str := "服务名称" + businessName + "\n告警名称:" + M.Name + "\n告警类型:" + M.Scope + "\n告警规则:" + M.RuleName + "" + "\n告警信息:" + M.AlarmMessage
	data["text"] = map[string]interface{}{
		"content": str,
	}
	data["safe"] = "0"
	Url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + Token
	//log.Println(str)
	HttpMethod := "POST"
	HttpBody, err := HttpRequest(HttpMethod, Url, data)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(HttpBody))
}
