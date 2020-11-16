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
	"errors"
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

type BuinessStruct struct {
	CRM []string `yaml:"CRM"`
	SSM []string `yaml:"SSM"`
	DLM []string `yaml:"DLM"`
	VX  []string `yaml:"VX"`
}

type TencentConfig struct {
	Auth         AuthStr       `yaml:"auth"`
	Agentid      string        `yaml:"agentid"`
	BusinessType BuinessStruct `yaml:"businesstype"`
	User         []string      `yaml:"user"`
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
		}
	}
	//fmt.Println(string(file))
	err = yaml.Unmarshal(file, &conf)
	//fmt.Printf("%#v\n", conf.Tencent.BusinessType)

	//fmt.Println("Bu",businessName)
	fmt.Printf("%+v", conf.Tencent.BusinessType)
	//for k,v:=conf.Tencent.BusinessType[]
	if err != nil {
		errors.New("config parase failed")
	}
}
