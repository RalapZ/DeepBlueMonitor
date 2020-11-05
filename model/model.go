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

package model

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
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

func (conf *Config) ReadConfig(filename string) {

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
