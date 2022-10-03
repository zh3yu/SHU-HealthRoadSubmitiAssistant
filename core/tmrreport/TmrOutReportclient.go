package tmrreport

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"regexp"
	"selfreport/cntime"
	. "selfreport/core/RedirectClient"
	"strings"
	"time"
)

type F_STATE_LXSQ struct {
	PersinfoCtl00 struct {
		IFrameAttributes struct {
		} `json:"IFrameAttributes"`
	} `json:"persinfo_ctl00"`
	PersinfoXiZhi struct {
		Checked bool `json:"Checked"`
	} `json:"persinfo_XiZhi"`
	PersinfoXueGH struct {
		Text string `json:"Text"`
	} `json:"persinfo_XueGH"`
	PersinfoXingMing struct {
		Text string `json:"Text"`
	} `json:"persinfo_XingMing"`
	PersinfoXueYBM struct {
		Text string `json:"Text"`
	} `json:"persinfo_XueYBM"`
	PersinfoSuiSMZT struct {
		Text string `json:"Text"`
	} `json:"persinfo_SuiSMZT"`
	PersinfoPHeSJCHeSJCLST struct {
		FItems [][]interface{} `json:"F_Items"`
	} `json:"persinfo_P_HeSJC_HeSJCLST"`
	PersinfoPHeSJCHeSYXSJ struct {
		Text string `json:"Text"`
	} `json:"persinfo_P_HeSJC_HeSYXSJ"`
	PersinfoPHeSJC struct {
		Title            string `json:"Title"`
		IFrameAttributes struct {
		} `json:"IFrameAttributes"`
	} `json:"persinfo_P_HeSJC"`
	PersinfoSuoZXQ struct {
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue string          `json:"SelectedValue"`
	} `json:"persinfo_SuoZXQ"`
	PersinfoHuanCQTip struct {
		IFrameAttributes struct {
		} `json:"IFrameAttributes"`
	} `json:"persinfo_HuanCQTip"`
	PersinfoChuXRQ struct {
		Readonly bool   `json:"Readonly"`
		Text     string `json:"Text"`
	} `json:"persinfo_ChuXRQ"`
	PersinfoYuanYin struct {
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue string          `json:"SelectedValue"`
	} `json:"persinfo_YuanYin"`
	PersinfoTeSYYQiTa struct {
		Hidden   bool   `json:"Hidden"`
		Required bool   `json:"Required"`
		Text     string `json:"Text"`
	} `json:"persinfo_TeSYY_QiTa"`
	PersinfoDdlSheng struct {
		FItems             [][]interface{} `json:"F_Items"`
		SelectedValueArray []string        `json:"SelectedValueArray"`
	} `json:"persinfo_ddlSheng"`
	PersinfoDdlShi struct {
		FItems             [][]interface{} `json:"F_Items"`
		SelectedValueArray []string        `json:"SelectedValueArray"`
		Enabled            bool            `json:"Enabled"`
	} `json:"persinfo_ddlShi"`
	PersinfoDdlXian struct {
		FItems             [][]interface{} `json:"F_Items"`
		SelectedValueArray []string        `json:"SelectedValueArray"`
		Enabled            bool            `json:"Enabled"`
	} `json:"persinfo_ddlXian"`
	PersinfoDangTHX struct {
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue interface{}     `json:"SelectedValue"`
	} `json:"persinfo_DangTHX"`
	PersinfoCtl01BtnSubmit struct {
		Hidden bool `json:"Hidden"`
	} `json:"persinfo_ctl01_btnSubmit"`
	Persinfo struct {
		IFrameAttributes struct {
		} `json:"IFrameAttributes"`
	} `json:"persinfo"`
}

type PostTmrReportInfo struct {
	CurrentCampus  string
	ViewState      string
	Fstatetemplate *F_STATE_LXSQ
	HeSJCinfo      *HeSJCinfo
}

func (s *PostTmrReportInfo) GetKvs() map[string]string {
	kvs := make(map[string]string)
	kvs["__EVENTTARGET"] = "persinfo$ctl01$btnSubmit"
	kvs["__EVENTARGUMENT"] = ""
	kvs["__VIEWSTATE"] = s.ViewState
	kvs["__VIEWSTATEGENERATOR"] = "5EBC3AEC"
	kvs["persinfo$XiZhi"] = "persinfo$XiZhi"
	kvs["persinfo$SuoZXQ"] = s.CurrentCampus + "校区"
	kvs["persinfo$ChuXRQ"] = cntime.NowCN().Add(24 * time.Hour).Format("2006-01-02")
	kvs["persinfo$YuanYin"] = "其他原因"
	kvs["persinfo$TeSYY_QiTa"] = "因事外出"
	kvs["persinfo$ddlSheng$Value"] = "上海"
	kvs["persinfo$ddlSheng"] = "上海"
	kvs["persinfo$ddlShi$Value"] = "上海市"
	kvs["persinfo$ddlShi"] = "上海市"
	kvs["persinfo$ddlXian$Value"] = "宝山区"
	kvs["persinfo$ddlXian"] = "宝山区"
	kvs["persinfo$XiangXDZ"] = "大场镇"
	kvs["persinfo$DangTHX"] = "是"
	kvs["persinfo_ctl00_Collapsed"] = "false"
	kvs["persinfo_P_HeSJC_Collapsed"] = "false"
	kvs["persinfo_HuanCQTip_Collapsed"] = "false"
	kvs["persinfo_Collapsed"] = "false"
	kvs["F_STATE"] = s.GetFstate()
	kvs["F_TARGET"] = "persinfo_ctl01_btnSubmit"
	return kvs
}
func (p *PostTmrReportInfo) GetFstate() string {
	fstate := p.Fstatetemplate
	// SuoZXQ 所在校区
	fstate.PersinfoSuoZXQ.SelectedValue = p.CurrentCampus
	// HeSJC_HeSJCLST
	var HeSJCLST [][]interface{}
	for _, HeSJCimte := range p.HeSJCinfo.HeSJC_HeSJCLST {
		HeSJCLST = append(HeSJCLST, []interface{}{"", HeSJCimte, 1, "", "", "", false})
	}
	fstate.PersinfoPHeSJCHeSJCLST.FItems = HeSJCLST
	// HeSJC_HeSJYXSJ
	fstate.PersinfoPHeSJCHeSYXSJ.Text = p.HeSJCinfo.HeSJC_HeSJYXSJ
	// HeSJC
	fstate.PersinfoPHeSJC.Title = p.HeSJCinfo.HeSJC
	// ChuXRQ 出校申请时间
	fstate.PersinfoChuXRQ.Text = cntime.NowCN().Add(24 * time.Hour).Format("2006-01-02")
	str, _ := json.Marshal(fstate)
	// fmt.Println(string(str))
	encodestate := base64.StdEncoding.EncodeToString(str)

	t := int(len(encodestate) / 2)
	// + "F_STATE" +
	encodestate = encodestate[:t] + "F_STATE" + encodestate[t:]
	return encodestate
}

type HeSJCinfo struct {
	HeSJC_HeSJCLST        []string
	HeSJC_HeSJYXSJ, HeSJC string
}

type TmrOutClient struct {
	Client *RedirectClient
}

func (s *TmrOutClient) CheakTmrOut() bool {
	t := cntime.NowCN().Add(24 * time.Hour)
	res, err := s.Client.BanRedirectGet("https://selfreport.shu.edu.cn/XiaoYJC202207/XueSLXSQ_List.aspx")
	if err != nil {
		fmt.Println("获取离校申请表失败")
		return false
	}
	body, _ := ioutil.ReadAll(res.Body)
	HTML := string(body)
	return strings.Contains(HTML, fmt.Sprintf("日期：%s", t.Format("2006-01-02")))
}
func (s *TmrOutClient) ReportTmrOut() error {
	if s.CheakTmrOut() {
		fmt.Println("当天的离校申请已提交，不需要再次申请")
		return nil
	}
	buf := new(bytes.Buffer)
	bw := multipart.NewWriter(buf)
	CurrentCampus, err := s.GetCurrentCampus()
	if err != nil {
		return err
	}
	view_state, HeSJCInfo, err := s.GetTmrViewstateAndHeSJCInfo()
	if err != nil {
		return err
	}
	PRI := &PostTmrReportInfo{CurrentCampus: CurrentCampus, ViewState: view_state, Fstatetemplate: s.GetFstatelxsqtemplate(), HeSJCinfo: HeSJCInfo}
	kvs := PRI.GetKvs()
	Addmultipartkvs(bw, kvs)
	bw.Close()
	req, _ := NewRequest("POST", "https://selfreport.shu.edu.cn/XiaoYJC202207/XueSLXSQ.aspx", buf)
	req.Header.Add("Content-Type", bw.FormDataContentType())
	req.Header.Add("X-FineUI-Ajax", "true")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	res, err := s.Client.BanRedirectDo(req)
	if err != nil {
		return err
	}
	body, _ := ioutil.ReadAll(res.Body)

	if strings.Contains(string(body), "提交成功") || strings.Contains(string(body), "该日期已申请，不可重复申请") || strings.Contains(string(body), "申请已提交") {
		fmt.Printf("离校申请Post成功\n")
		return nil
	} else {
		fmt.Printf("离校申请Post失败\n")
		fmt.Println(string(body))
		return errors.New(fmt.Sprintf("离校申请Post失败\n"))
	}

}
func (s *TmrOutClient) GetFstatelxsqtemplate() *F_STATE_LXSQ {
	fstate := new(F_STATE_LXSQ)
	fstateFile, err := ioutil.ReadFile("Resources/fstate_LXSQ.json")
	if err != nil {
		panic("读取离校申请模板错误")
	}
	json.Unmarshal(fstateFile, fstate)
	return fstate
}
func (s *TmrOutClient) GetCurrentCampus() (string, error) {
	t := cntime.NowCN().Add(-24 * time.Hour)
	url := fmt.Sprintf("https://selfreport.shu.edu.cn/ViewDayReport.aspx?day=%s", t.Format("2006-01-02"))
	var re = regexp.MustCompile(`(?m)XiaoQu.*<span>(宝山)<\/span>'}`)
	req, _ := NewRequest("GET", url, nil)
	res, err := s.Client.Do(req)
	if err != nil {
		return "", err
	}
	body, _ := ioutil.ReadAll(res.Body)
	match := re.FindStringSubmatch(string(body))
	if match == nil {
		return "", errors.New("页面无校区数据")
	}
	return match[1], nil
}
func (s *TmrOutClient) GetTmrViewstateAndHeSJCInfo() (string, *HeSJCinfo, error) {
	url := "https://selfreport.shu.edu.cn/XiaoYJC202207/XueSLXSQ.aspx"
	rep, err := s.Client.Get(url)
	if err != nil {
		return "", &HeSJCinfo{}, fmt.Errorf("网络连接错误:%v", err)
	}
	body, _ := ioutil.ReadAll(rep.Body)
	if strings.Contains(string(body), "message:'申请时间为8:00-17:00'") {
		// 	申请时间为8:00-17:00
		return "", &HeSJCinfo{}, errors.New("不在申请时间内")
	}
	viewstate := s.getTmrViewState(string(body))
	heSJCinfo, err := s.getTmrHeSJC(string(body))
	if err != nil {
		return "", &HeSJCinfo{}, fmt.Errorf("页面错误:%v", err)
	}
	return viewstate, heSJCinfo, nil
}
func (s *TmrOutClient) getTmrViewState(body string) string {
	var retstr string
	var re = regexp.MustCompile(`(?mU)id="__VIEWSTATE" value="([\S\s]*)" \/>`)
	match := re.FindStringSubmatch(string(body))
	if len(match) == 0 {
		fmt.Printf("没有找到ViewState")
		return ""
	}
	retstr = match[1]
	return retstr
}
func (s *TmrOutClient) getTmrHeSJC(body string) (*HeSJCinfo, error) {
	HeSJC_HeSJCLSTre := regexp.MustCompile(`(?mU)"(采样时间：.*<br\/>报告时间：.*<br\/>检测结果：.*[^\[,])"`)
	HeSJC_HeSYXSJre := regexp.MustCompile(`(?mU)"(根据当前核酸检测结果，可在.*前进校)"`)
	HeSJCre := regexp.MustCompile(`(?mU)(最近的核酸检测情况<br/>（大数据中心同步时间为.*）)`)
	var matchlist1 [][]string = HeSJC_HeSJCLSTre.FindAllStringSubmatch(string(body), -1)
	var match2 []string = HeSJC_HeSYXSJre.FindStringSubmatch(string(body))
	var match3 []string = HeSJCre.FindStringSubmatch(string(body))
	var HeSJCLST []string
	for _, match := range matchlist1 {
		if len(match) > 1 {
			HeSJCLST = append(HeSJCLST, match[1])
		}
	}
	if len(matchlist1) == 0 || len(match2) == 0 || len(match3) == 0 {
		fmt.Println("没有核酸检测信息")
		return &HeSJCinfo{}, errors.New("没有核酸检测信息")
	}
	fmt.Printf(`核酸检测信息：
		%v
		%v
		%v
	`, HeSJCLST, match2[1], match3[1])
	HeSJCinfo := &HeSJCinfo{HeSJC_HeSJCLST: HeSJCLST, HeSJC_HeSJYXSJ: match2[1], HeSJC: match3[1]}
	return HeSJCinfo, nil
}
