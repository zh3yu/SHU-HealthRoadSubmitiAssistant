package tmrreport_test

import (
	"fmt"
	"os"
	"selfreport/core"
	"selfreport/core/login"
	"selfreport/core/message"
	"selfreport/core/tmrreport"
	"testing"
)

var tmrlient *tmrreport.TmrOutClient

func TestMain(m *testing.M) {
	os.Chdir("../../")
	usc := &login.UserClient{}
	usc.Init(core.Name, core.PassWord)
	client, _ := usc.GetLoginedClient()
	msgclient := &message.Messageclient{Client: client}
	msgclient.VisitUnreadMessage()
	tmrlient = &tmrreport.TmrOutClient{Client: client}
	retcode := m.Run()
	os.Exit(retcode)
}

func TestCheakTmrOut(t *testing.T) {
	println(tmrlient.CheakTmrOut())
}

func TestPostTMR(t *testing.T) {
	err := tmrlient.ReportTmrOut()
	if err != nil {
		t.Error(err)
	}
}

func TestGetCurrentCampus(t *testing.T) {
	ShuoZXQ, err := tmrlient.GetCurrentCampus()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ShuoZXQ)
}
func TestGetViewAndHSJC(t *testing.T) {
	tmrlient.GetTmrViewstateAndHeSJCInfo()

}

func TestGetfstate(t *testing.T) {
	CurrentCampus, err := tmrlient.GetCurrentCampus()
	if err != nil {
		t.Error(err)
	}
	view_state, HeSJCInfo, err := tmrlient.GetTmrViewstateAndHeSJCInfo()
	if err != nil {
		t.Error(err)
	}
	PRI := &tmrreport.PostTmrReportInfo{CurrentCampus: CurrentCampus, ViewState: view_state, Fstatetemplate: tmrlient.GetFstatelxsqtemplate(), HeSJCinfo: HeSJCInfo}
	fmt.Println(PRI.GetFstate())
}
