package selfreport_test

import (
	"log"
	"os"
	"selfreport/cntime"
	"selfreport/core"
	"selfreport/core/login"
	"selfreport/core/message"
	"selfreport/core/selfreport"
	"testing"
	"time"
)

var srclient *selfreport.SelfReportClient

func TestMain(m *testing.M) {
	os.Chdir("../../")
	usc := &login.UserClient{}
	usc.Init(core.Name, core.PassWord) //请填入学号，密码
	client, _ := usc.GetLoginedClient()
	msgclient := &message.Messageclient{Client: client}
	msgclient.VisitUnreadMessage()
	srclient = &selfreport.SelfReportClient{Client: client}
	retcode := m.Run()
	os.Exit(retcode)
}

func TestReport(t *testing.T) {
	srclient.Report(cntime.NowCN())
}
func TestPostReport(t *testing.T) {
	// coreInfo := &CoreInfo{ShiFSH: "在上海（校内）", JinXXQ: "宝山", ShiFZX: "是",
	// 	XiaoQu: "宝山", ddlSheng: "上海", ddlShi: "上海市", ddlXian: "宝山区", ddlJieDao: "大场镇", XiangXDZ: "上海大学校内", ShiFZJ: "否"}
	// coreInfo = srclient.ParseReportInfo(srclient.GetReportInfo(NowCN()))
	s := srclient.GetReportInfo(cntime.NowCN().Add(-24 * time.Hour))
	s2 := srclient.ParseReportInfo(s)
	log.Println(s2)
	log.Println(srclient.PostReport(s2, cntime.NowCN()))

}

func TestReportinfo(t *testing.T) {
	body := srclient.GetReportInfo(cntime.NowCN().Add(-60 * time.Hour * 24))
	log.Println(srclient.ParseReportInfo(body))
	body = srclient.GetReportInfo(cntime.NowCN().Add(-15 * time.Hour * 24))
	log.Println(srclient.ParseReportInfo(body))

}

func TestGetPhoneNum(t *testing.T) {
	log.Println(srclient.GetPhoneNum())
}

func TestGetViewState(t *testing.T) {
	log.Println(srclient.GetViewState(cntime.NowCN()))
}
func TestGetFstateT(t *testing.T) {
	srclient.GetFstatedaytemplate()
}

var array = []int{10, 9, 10, 10, 3, 4, 5, 0, 0, 0}

func LogicReport(index int) int {
	if cheakInfo(index) {
		return array[index]
	} else {
		lastint := LogicReport(index - 1)
		array[index] = lastint
		return array[index]
	}
}
func cheakInfo(index int) bool {
	if array[index] == 0 {
		return false
	} else {
		return true
	}

}
func TestLogicReport(t *testing.T) {
	log.Println(LogicReport(9), array)
	//5 [10 9 10 10 3 4 5 5 5 5]
}
