package login_test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"selfreport/core"
	"selfreport/core/login"
	"testing"
)

var user *login.User
var usc *login.UserClient

func TestMain(m *testing.M) {
	os.Chdir("../../")
	usc := &login.UserClient{}
	usc.Init(core.Name, core.PassWord)
	ret := m.Run()
	os.Exit(ret)
}
func TestGet(t *testing.T) {

	UserFactory := &login.UserFactory{}
	user = UserFactory.Get("12345678", "password")
	fmt.Println(user.Password)
	if user.Password[len(user.Password)-1] != '=' {
		t.Errorf("'错误'")
	}
}

func TestJsonMarshal(t *testing.T) {
	bodyBytes, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(bodyBytes))

}
func TestLogin(t *testing.T) {

	_, err := usc.GetLoginedClient()
	if err != nil {
		log.Fatal(err)
	}

}
