package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"selfreport/cntime"
	"selfreport/core/login"
	"selfreport/core/message"
	"selfreport/core/selfreport"
	"selfreport/core/tmrreport"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type CONF struct {
	Users []struct {
		Name     string `yaml:"name"`
		PassWord string `yaml:"password"`
	} `yaml:"Users"`
}

func main() {
	fmt.Println(os.Getwd())
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	if !strings.Contains(ex, "exe\\main.exe") {
		os.Chdir(exPath)
	}
	fmt.Println(`本项目仅作为免费的网络研究使用，
不得利用本程序以任何方式直接或者间接的从事违反中国法律、国际公约以及社会公德的行为，
！！！不支持进行虚假填报！！！
	`)
	conf, err := ioutil.ReadFile("Configuration/AccountList.yaml")
	if err != nil {
		panic(fmt.Errorf("配置文件打开错误 %v", err))
	}
	cntime.PrintNow()
	var CONF = new(CONF)
	err = yaml.Unmarshal(conf, CONF)
	if err != nil {
		panic(fmt.Errorf("yaml解析错误 %v", err))
	}
	for _, user := range CONF.Users {
		fmt.Println(user.Name)
		usc := &login.UserClient{}
		usc.Init(user.Name, user.PassWord) //请填入学号，密码
		client, err := usc.GetLoginedClient()
		if err != nil {
			fmt.Printf("学号%v 登陆错误", user.Name)
			fmt.Println(err)
			continue
		}
		msgclient := &message.Messageclient{Client: client}
		msgclient.VisitUnreadMessage()
		srclient := &selfreport.SelfReportClient{Client: client}
		// 每日一报
		_, err = srclient.Report(cntime.NowCN())
		if err != nil {
			fmt.Printf("学号%v 每日一报错误", user.Name)
			fmt.Println(err)
		}
		// 离校申请
		tmrlient := &tmrreport.TmrOutClient{Client: client}
		err = tmrlient.ReportTmrOut()
		if err != nil {
			fmt.Printf("学号%v 离校申请错误", user.Name)
			fmt.Println(err)
		}

	}

	time.Sleep(10 * time.Second)

}
