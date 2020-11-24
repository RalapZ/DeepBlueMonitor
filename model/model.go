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
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"time"
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

type BuinessInfo struct {
	BusinesType string
	BusinesName string
}

type SkywalkInfo struct {
	ScopeId       int    `json:scopeId`
	Scope         string `json:scope`
	Name          string `json:name`
	Id0           string `json:id0`
	Id1           string `json:id1`
	RuleName      string `json:ruleName`
	AlarmMessage  string `json:alarmMessage`
	StartTime     time.Time
	Starttimerecv int64 `json:"startTime"`
	BInfo         BuinessInfo
}

var CONF *Config

func (conf *Config) ReadConfig(filename string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if error, ok := err.(*os.PathError); ok {
			log.Println(error.Op, "file not exist", error.Err)
			panic(error.Error())
		}
	}
	err = yaml.Unmarshal(file, &conf)
	CONF = conf
	//fmt.Printf("%+v", conf.Tencent.BusinessType)
	if err != nil {
		errors.New("config parase failed")
	}

}
