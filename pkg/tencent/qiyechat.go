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
//   Date  : 2020/09/17                                                                             //
//##################################################################################################//
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

type AuthStr struct {
	CorpInfo   string `yaml:"corpid"`
	CorpSecret string `yaml:"corpsecret"`
}

type TencentConfig struct {
	Auth    AuthStr  `yaml:"auth"`
	Agentid string   `yaml:"agentid"`
	User    []string `yaml:"user"`
}

type Config struct {
	Tencent    TencentConfig `yaml:"tencent"`
	Listenport string        `yaml:"listenport"`
}

type CorpInfo struct {
	corpid     string
	corpsecret string
}
type QiYchatStu struct {
	Errcod       int    `json:"errocd"`
	Errmsg       string `json:errmsg`
	Access_token string `json:access_token`
	expires_in   int
}

type SkywalkInfo struct {
	ScopeId      int    `json:scopeId`
	Scope        string `json:scope`
	Name         string `json:name`
	id0          string
	id1          string
	RuleName     string `json:ruleName`
	AlarmMessage string `json:alarmMessage`
	StartTime    string `json:startTime`
}

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
	qistu := QiYchatStu{}
	_ = json.Unmarshal(body, &qistu)
	return qistu.Access_token
}

//func SendMessage(CorpId string, CorpSecret string, M SkywalkInfo,conf Config) {
func SendMessage(M SkywalkInfo, conf Config) {
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

func (conf *Config) ReadConfig(filename string) {
	//filename:="src/project/SkywalkingWebHook/conf/application.yaml1"
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if error, ok := err.(*os.PathError); ok {
			log.Println(error.Op, "file not exist", error.Err)
			panic(error.Error())
			//os.Exit(0)
		}
	}
	fmt.Println(string(file))
	//fmt.Println(yaml.Unmarshal(file, &conf))
	yaml.Unmarshal(file, &conf)

}

func MainFunc(conf Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		test, _ := ioutil.ReadAll(r.Body)
		var DataInfo []SkywalkInfo
		json.Unmarshal(test, &DataInfo)
		//fmt.Println(&conf)
		for _, Message := range DataInfo {
			//corpid := "ww97af1eab5d2add3c"
			//corpsecret := "iq2IyxRcY3oCHHTFg2U2o3UQGHzXIWkKIAgfKQFdhxw"
			//SendMessage(corpid, corpsecret, Message,conf)
			SendMessage(Message, conf)
		}
	}
}

func main() {
	filename := "C:/code/DeepBlueMonitor/conf/application.yaml"
	var conf Config
	//str,_:=os.Getwd()
	//fmt.Println(string(str))
	conf.ReadConfig(filename)
	//fmt.Println(conf)
	//fmt.Println(conf.Listenport)
	http.HandleFunc("/alarm", MainFunc(conf))
	fmt.Println(conf)
	http.ListenAndServe(fmt.Sprintf(":%s", conf.Listenport), nil)
}
