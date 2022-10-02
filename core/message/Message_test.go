package message_test

import (
	"log"
	"os"
	"selfreport/core"
	"selfreport/core/login"
	"selfreport/core/message"
	"testing"
)

func TestGetMessageHtml(t *testing.T) {
	os.Chdir("../../")
	usc := &login.UserClient{}
	usc.Init(core.Name, core.PassWord)
	client, err := usc.GetLoginedClient()
	if err != nil {
		log.Fatal(err)
	}
	msgclient := &message.Messageclient{Client: client}
	msgclient.VisitUnreadMessage()
}
