package login

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	url0 "net/url"
	. "selfreport/core/RedirectClient"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type User struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

type UserFactory struct{}

func (u *UserFactory) encryptPass(password string) (string, error) {
	var publicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDl/aCgRl9f/4ON9MewoVnV58OL
OU2ALBi2FKc5yIsfSpivKxe7A6FitJjHva3WpM7gvVOinMehp6if2UNIkbaN+plW
f5IwqEVxsNZpeixc4GsbY9dXEk3WtRjwGSyDLySzEESH/kpJVoxO7ijRYqU+2oSR
wTBNePOk1H+LRQokgQIDAQAB
-----END PUBLIC KEY-----`)
	block, _ := pem.Decode(publicKey)
	if block == nil {
		panic("解码错误")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	pub := pubInterface.(*rsa.PublicKey)
	rsa1, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(password))
	ret := base64.StdEncoding.EncodeToString(rsa1)
	return ret, err
}

func (u *UserFactory) Get(name string, password string) *User {
	encryptedPass, err := u.encryptPass(password)
	if err != nil {
		panic("res加密失败")
	}
	return &User{
		Name:     name,
		Password: encryptedPass,
	}
}

type Loginclientstate struct {
	Timestamp    int64  `json:"timestamp"`
	ResponseType string `json:"responseType"`
	ClientID     string `json:"clientId"`
	Scope        string `json:"scope"`
	RedirectURI  string `json:"redirectUri"`
	State        string `json:"state"`
}

type UserClient struct {
	user     *User
	client   *RedirectClient
	logincst *Loginclientstate
	LastRep  *http.Response
}

func (u *UserClient) Init(name string, password string) {
	UserFactory := &UserFactory{}
	user := UserFactory.Get(name, password)
	u.user = user
	jar, _ := cookiejar.New(nil)
	u.client = &RedirectClient{&http.Client{Jar: jar}}
}

func (c *UserClient) getLoginclientstate() (*Loginclientstate, error) {
	url := "https://selfreport.shu.edu.cn/Default.aspx"
	req, err := NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	rep, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	c.LastRep = rep
	resurlt := rep.Request.URL.Path
	resSplit := strings.Split(resurlt, "/")
	code := resSplit[len(resSplit)-1]
	url_param_json, err := base64.URLEncoding.DecodeString(code)
	if err != nil {
		return nil, err
	}
	logincst := new(Loginclientstate)
	err = json.Unmarshal([]byte(url_param_json), logincst)
	if err != nil {
		return nil, err
	}
	return logincst, nil
}

func (c *UserClient) LoginPost() error {
	logincst, err := c.getLoginclientstate()
	if err != nil {
		panic(err)
	}
	c.logincst = logincst

	c.client.SetCheckRedirect(false)
	defer c.client.SetCheckRedirect(true)

	url := c.LastRep.Request.URL.String()
	bodytemplate := "username=%s&password=%s"
	// str := "XiQZ%2B7N82A%2Fxnzuqedt%2BuGNS8ymnJlo2Zm%2Bg4GOp1aAlN6NwyxAcQPKxuYvRyHJFgSOe%2B2kAxoMtzHjkELZePn%2BwsU8Sx46pegLyfGaKtwZRuTa2vSq6mvmy1yB90hgQEolExkBmqDuA3Oq3lw%2FJ0wLMMa5%2FgixqZp42ZAk%2Ba6o%3D"
	rawstr := "XiQZ+7N82A/xnzuqedt+uGNS8ymnJlo2Zm+g4GOp1aAlN6NwyxAcQPKxuYvRyHJFgSOe+2kAxoMtzHjkELZePn+wsU8Sx46pegLyfGaKtwZRuTa2vSq6mvmy1yB90hgQEolExkBmqDuA3Oq3lw/J0wLMMa5/gixqZp42ZAk+a6o="
	rawstr = c.user.Password
	body := fmt.Sprintf(bodytemplate, url0.QueryEscape(c.user.Name), url0.QueryEscape(rawstr))
	req, _ := NewRequest("POST", url, strings.NewReader(body))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rep, err := c.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		log.Fatal(err)
	}

	if find := strings.Contains(string(data), "上海大学统一身份认证"); find {
		return errors.New(" ：登陆失败，未能跳转到正确页面")
	}
	fmt.Println(("登陆成功"))
	c.LastRep = rep
	return nil
}
func (c *UserClient) RedirectHome() error {
	url := "https://newsso.shu.edu.cn/oauth/authorize?response_type=code&client_id=WUHWfrntnWYHZfzQ5QvXUCVy&redirect_uri=https%3a%2f%2fselfreport.shu.edu.cn%2fLoginSSO.aspx%3fReturnUrl%3d%252fDefault.aspx&scope=1&" + fmt.Sprintf("state=%s", c.logincst.State)
	req, err := NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	_, err = c.client.Do(req)
	if err != nil {
		panic(err)
	}
	return nil
}

func (c *UserClient) testHome() error {
	req, _ := NewRequest("GET", "https://selfreport.shu.edu.cn/DayReport.aspx", nil)
	res, err := c.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	doc, err := goquery.NewDocumentFromReader(res.Body)
	dom := doc.Find("input [name: '__VIEWSTATE']")
	if len(dom.Nodes) != 0 {
		return nil
	}
	return errors.New("没有正确登陆")
}
func (c *UserClient) GetLoginedClient() (*RedirectClient, error) {

	err := c.LoginPost()
	if err != nil {
		return nil, err
	}
	err = c.RedirectHome()
	if err != nil {
		return nil, err
	}
	// err = c.testHome()
	// if err != nil {
	// 	return nil, err
	// }
	return c.client, nil
}
