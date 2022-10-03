package message

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	. "selfreport/core/RedirectClient"
	"strings"
)

type Messageclient struct {
	Client      *RedirectClient
	MessageHtml string
}

type FItemss struct {
	FItems [][]interface{} `json:"F_Items"`
}

func (m *Messageclient) InitMessageHtml() error {
	client := m.Client
	req, err := NewRequest("GET", "https://selfreport.shu.edu.cn/MyMessages.aspx", nil)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(res.Body)
	var build strings.Builder
	for {
		line, _, err := reader.ReadLine()
		if err != nil || io.EOF == err {
			break
		}
		build.Write(line)
	}
	m.MessageHtml = build.String()
	if err != nil {
		log.Println("获得消息错误")
		return err
	}
	return nil
}

func (m *Messageclient) getUrlList() []string {
	str := m.MessageHtml
	var re = regexp.MustCompile(`var .._state=({"F_Items":\[\[[\S\s]*\]\]});`)
	for _, match := range re.FindAllStringSubmatch(str, -1) {
		str = match[1]
	}
	FItemss := new(FItemss)
	ERR := json.Unmarshal([]byte(str), FItemss)
	_ = ERR
	var strlists []string
	for _, items := range FItemss.FItems {
		var stface []string
		for i, _ := range items {
			str1 := fmt.Sprintf("%v", items[i])
			stface = append(stface, str1)
		}
		strlist := strings.Join(stface, ",")
		// 未实际运行 使用 strings.Contains(strlist, "标题") 做测试
		if strings.Contains(strlist, "未读") {
			str := fmt.Sprintf("https://selfreport.shu.edu.cn%v", items[4])
			strlists = append(strlists, str)
		}
	}
	return strlists
}
func (m *Messageclient) VisitUrl(urls []string) {
	log.Println("未读的链接为", urls)
	for _, url := range urls {
		rep, _ := m.Client.BanRedirectGet(url)
		ioutil.ReadAll(rep.Body)
		log.Println(url, "已读")
	}
}
func (msgclient *Messageclient) VisitUnreadMessage() {
	msgclient.InitMessageHtml()
	urls := msgclient.getUrlList()
	msgclient.VisitUrl(urls)

}
