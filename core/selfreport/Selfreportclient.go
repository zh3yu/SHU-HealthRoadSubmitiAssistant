package selfreport

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"regexp"
	"selfreport/cntime"
	. "selfreport/core/RedirectClient"
	"strings"
	"time"

	"github.com/fogleman/gg"
)

type F_STATE_IMAGE struct {
	P1Tip struct {
		IFrameAttributes struct {
		} `json:"IFrameAttributes"`
	} `json:"p1_Tip"`
	P1PnlDangSZSDangSZS struct {
		FItems             []interface{} `json:"F_Items"`
		SelectedValueArray []interface{} `json:"SelectedValueArray"`
	} `json:"p1_pnlDangSZS_DangSZS"`
	P1PnlDangSZS struct {
		IFrameAttributes struct {
		} `json:"IFrameAttributes"`
	} `json:"p1_pnlDangSZS"`
	P1PQueZXXCengQZ struct {
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue interface{}     `json:"SelectedValue"`
	} `json:"p1_P_QueZXX_CengQZ"`
	P1PQueZXXQueZDSheng struct {
		FItems             []interface{} `json:"F_Items"`
		SelectedValueArray []interface{} `json:"SelectedValueArray"`
	} `json:"p1_P_QueZXX_QueZD_Sheng"`
	P1PQueZXXQueZDShi struct {
		FItems             []interface{} `json:"F_Items"`
		SelectedValueArray []interface{} `json:"SelectedValueArray"`
	} `json:"p1_P_QueZXX_QueZD_Shi"`
	P1PQueZXX struct {
		IFrameAttributes struct {
		} `json:"IFrameAttributes"`
	} `json:"p1_P_QueZXX"`
	P1BaoSRQ struct {
		Text string `json:"Text"`
	} `json:"p1_BaoSRQ"`
	P1CengFWSS struct {
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue interface{}     `json:"SelectedValue"`
	} `json:"p1_CengFWSS"`
	P1TuJing202210 struct {
		FItems             [][]interface{} `json:"F_Items"`
		SelectedValueArray []interface{}   `json:"SelectedValueArray"`
	} `json:"p1_TuJing_202210"`
	P1DangQSTZK struct {
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue string          `json:"SelectedValue"`
	} `json:"p1_DangQSTZK"`
	P1ZhengZhuang struct {
		Hidden             bool            `json:"Hidden"`
		FItems             [][]interface{} `json:"F_Items"`
		SelectedValueArray []interface{}   `json:"SelectedValueArray"`
	} `json:"p1_ZhengZhuang"`
	P1GuoNei struct {
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue string          `json:"SelectedValue"`
	} `json:"p1_GuoNei"`
	P1PGuoNeiShiFSH struct {
		Hidden        bool            `json:"Hidden"`
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue interface{}     `json:"SelectedValue"`
	} `json:"p1_P_GuoNei_ShiFSH"`
	P1PGuoNeiJinXXQ struct {
		FItems             [][]interface{} `json:"F_Items"`
		SelectedValueArray []interface{}   `json:"SelectedValueArray"`
	} `json:"p1_P_GuoNei_JinXXQ"`
	P1PGuoNeiShiFZX struct {
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue interface{}     `json:"SelectedValue"`
	} `json:"p1_P_GuoNei_ShiFZX"`
	P1PGuoNeiXiaoQu struct {
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue interface{}     `json:"SelectedValue"`
	} `json:"p1_P_GuoNei_XiaoQu"`
	P1PGuoNeiPImages struct {
		Hidden           bool `json:"Hidden"`
		IFrameAttributes struct {
		} `json:"IFrameAttributes"`
	} `json:"p1_P_GuoNei_pImages"`
	P1PGuoNei struct {
		Hidden           bool `json:"Hidden"`
		IFrameAttributes struct {
		} `json:"IFrameAttributes"`
	} `json:"p1_P_GuoNei"`
	P1JinChuSQ struct {
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue interface{}     `json:"SelectedValue"`
	} `json:"p1_JinChuSQ"`
	P1QiuZZT struct {
		FItems             []interface{} `json:"F_Items"`
		SelectedValueArray []interface{} `json:"SelectedValueArray"`
	} `json:"p1_QiuZZT"`
	P1JiuYKN struct {
		FItems             []interface{} `json:"F_Items"`
		SelectedValueArray []interface{} `json:"SelectedValueArray"`
	} `json:"p1_JiuYKN"`
	P1JiuYYX struct {
		Required           bool          `json:"Required"`
		FItems             []interface{} `json:"F_Items"`
		SelectedValueArray []interface{} `json:"SelectedValueArray"`
	} `json:"p1_JiuYYX"`
	P1JiuYZD struct {
		FItems             []interface{} `json:"F_Items"`
		SelectedValueArray []interface{} `json:"SelectedValueArray"`
	} `json:"p1_JiuYZD"`
	P1JiuYZL struct {
		FItems             []interface{} `json:"F_Items"`
		SelectedValueArray []interface{} `json:"SelectedValueArray"`
	} `json:"p1_JiuYZL"`
	P1DdlGuoJia struct {
		DataTextField      string          `json:"DataTextField"`
		DataValueField     string          `json:"DataValueField"`
		FItems             [][]interface{} `json:"F_Items"`
		SelectedValueArray []string        `json:"SelectedValueArray"`
	} `json:"p1_ddlGuoJia"`
	P1DdlSheng struct {
		FItems             [][]interface{} `json:"F_Items"`
		SelectedValueArray []string        `json:"SelectedValueArray"`
	} `json:"p1_ddlSheng"`
	P1DdlShi struct {
		Enabled            bool            `json:"Enabled"`
		FItems             [][]interface{} `json:"F_Items"`
		SelectedValueArray []string        `json:"SelectedValueArray"`
	} `json:"p1_ddlShi"`
	P1DdlXian struct {
		Enabled            bool            `json:"Enabled"`
		FItems             [][]interface{} `json:"F_Items"`
		SelectedValueArray []string        `json:"SelectedValueArray"`
	} `json:"p1_ddlXian"`
	P1DdlJieDao struct {
		Hidden             bool            `json:"Hidden"`
		FItems             [][]interface{} `json:"F_Items"`
		SelectedValueArray []string        `json:"SelectedValueArray"`
	} `json:"p1_ddlJieDao"`
	P1XiangXDZ struct {
		Text string `json:"Text"`
	} `json:"p1_XiangXDZ"`
	P1ShiFZJ struct {
		Required      bool            `json:"Required"`
		Hidden        bool            `json:"Hidden"`
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue interface{}     `json:"SelectedValue"`
	} `json:"p1_ShiFZJ"`
	P1GaoZDFXLJS struct {
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue interface{}     `json:"SelectedValue"`
	} `json:"p1_GaoZDFXLJS"`
	P1QueZHZJC struct {
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue interface{}     `json:"SelectedValue"`
	} `json:"p1_QueZHZJC"`
	P1DangRGL struct {
		SelectedValue string          `json:"SelectedValue"`
		FItems        [][]interface{} `json:"F_Items"`
	} `json:"p1_DangRGL"`
	P1GeLSM struct {
		Hidden           bool `json:"Hidden"`
		IFrameAttributes struct {
		} `json:"IFrameAttributes"`
	} `json:"p1_GeLSM"`
	P1GeLFS struct {
		Required      bool            `json:"Required"`
		Hidden        bool            `json:"Hidden"`
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue interface{}     `json:"SelectedValue"`
	} `json:"p1_GeLFS"`
	P1GeLDZ struct {
		Hidden bool `json:"Hidden"`
	} `json:"p1_GeLDZ"`
	P1Ctl01BtnReturn struct {
		OnClientClick string `json:"OnClientClick"`
	} `json:"p1_ctl01_btnReturn"`
	P1 struct {
		Title            string `json:"Title"`
		IFrameAttributes struct {
		} `json:"IFrameAttributes"`
	} `json:"p1"`
}

type Imageinfo struct {
	ViewState           string
	Time                time.Time
	Fstateimagetemplate *F_STATE_IMAGE
}

func (p *Imageinfo) GetFstate() string {
	fstate := p.Fstateimagetemplate
	// fstate['p1_BaoSRQ']['Text'] = BaoSRQ
	fstate.P1BaoSRQ.Text = p.Time.Format("2006-01-02")
	str, _ := json.Marshal(fstate)
	// log.Println(string(str))
	encodestate := base64.StdEncoding.EncodeToString(str)
	t := int(len(encodestate) / 2)
	// + "F_STATE" +
	encodestate = encodestate[:t] + encodestate[t:]
	return encodestate
}

type F_STATE_DAY struct {
	P1ChengNuo struct {
		Checked bool `json:"Checked"`
	} `json:"p1_ChengNuo"`
	P1PnlDangSZSDangSZS struct {
		FItems             []interface{} `json:"F_Items"`
		SelectedValueArray []interface{} `json:"SelectedValueArray"`
	} `json:"p1_pnlDangSZS_DangSZS"`
	P1PnlDangSZS struct {
		IFrameAttributes struct {
		} `json:"IFrameAttributes"`
	} `json:"p1_pnlDangSZS"`
	P1BaoSRQ struct {
		Text string `json:"Text"`
	} `json:"p1_BaoSRQ"`
	P1CengFWSS struct {
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue string          `json:"SelectedValue"`
	} `json:"p1_CengFWSS"`
	P1DangQSTZK struct {
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue string          `json:"SelectedValue"`
	} `json:"p1_DangQSTZK"`
	P1ZhengZhuang struct {
		Hidden             bool            `json:"Hidden"`
		FItems             [][]interface{} `json:"F_Items"`
		SelectedValueArray []interface{}   `json:"SelectedValueArray"`
	} `json:"p1_ZhengZhuang"`
	P1GuoNei struct {
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue string          `json:"SelectedValue"`
	} `json:"p1_GuoNei"`
	P1PGuoNeiShiFSH struct {
		Required      bool            `json:"Required"`
		Hidden        bool            `json:"Hidden"`
		SelectedValue string          `json:"SelectedValue"`
		FItems        [][]interface{} `json:"F_Items"`
	} `json:"p1_P_GuoNei_ShiFSH"`
	P1PGuoNeiJinXXQ struct {
		Required           bool            `json:"Required"`
		Hidden             bool            `json:"Hidden"`
		SelectedValueArray []string        `json:"SelectedValueArray"`
		FItems             [][]interface{} `json:"F_Items"`
	} `json:"p1_P_GuoNei_JinXXQ"`
	P1PGuoNeiShiFZX struct {
		Required      bool            `json:"Required"`
		Hidden        bool            `json:"Hidden"`
		SelectedValue string          `json:"SelectedValue"`
		FItems        [][]interface{} `json:"F_Items"`
	} `json:"p1_P_GuoNei_ShiFZX"`
	P1PGuoNeiXiaoQu struct {
		Hidden        bool            `json:"Hidden"`
		SelectedValue string          `json:"SelectedValue"`
		FItems        [][]interface{} `json:"F_Items"`
	} `json:"p1_P_GuoNei_XiaoQu"`
	P1PGuoNeiPImagesHFimgXingCM struct {
		Text string `json:"Text"`
	} `json:"p1_P_GuoNei_pImages_HFimgXingCM"`
	P1PGuoNeiPImagesImgXingCM struct {
		ImageURL string `json:"ImageUrl"`
	} `json:"p1_P_GuoNei_pImages_imgXingCM"`
	P1PGuoNeiPImages struct {
		Hidden           bool `json:"Hidden"`
		IFrameAttributes struct {
		} `json:"IFrameAttributes"`
	} `json:"p1_P_GuoNei_pImages"`
	P1PGuoNei struct {
		Hidden           bool `json:"Hidden"`
		IFrameAttributes struct {
		} `json:"IFrameAttributes"`
	} `json:"p1_P_GuoNei"`
	P1JinChuSQ struct {
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue interface{}     `json:"SelectedValue"`
	} `json:"p1_JinChuSQ"`
	P1QiuZZT struct {
		FItems             []interface{} `json:"F_Items"`
		SelectedValueArray []interface{} `json:"SelectedValueArray"`
	} `json:"p1_QiuZZT"`
	P1JiuYKN struct {
		FItems             []interface{} `json:"F_Items"`
		SelectedValueArray []interface{} `json:"SelectedValueArray"`
	} `json:"p1_JiuYKN"`
	P1JiuYYX struct {
		Required           bool          `json:"Required"`
		FItems             []interface{} `json:"F_Items"`
		SelectedValueArray []interface{} `json:"SelectedValueArray"`
	} `json:"p1_JiuYYX"`
	P1JiuYZD struct {
		FItems             []interface{} `json:"F_Items"`
		SelectedValueArray []interface{} `json:"SelectedValueArray"`
	} `json:"p1_JiuYZD"`
	P1JiuYZL struct {
		FItems             []interface{} `json:"F_Items"`
		SelectedValueArray []interface{} `json:"SelectedValueArray"`
	} `json:"p1_JiuYZL"`
	P1DdlGuoJia struct {
		DataTextField      string          `json:"DataTextField"`
		DataValueField     string          `json:"DataValueField"`
		FItems             [][]interface{} `json:"F_Items"`
		SelectedValueArray []string        `json:"SelectedValueArray"`
	} `json:"p1_ddlGuoJia"`
	P1DdlSheng struct {
		Hidden             bool            `json:"Hidden"`
		Readonly           bool            `json:"Readonly"`
		FItems             [][]interface{} `json:"F_Items"`
		SelectedValueArray []string        `json:"SelectedValueArray"`
	} `json:"p1_ddlSheng"`
	P1DdlShi struct {
		Hidden             bool            `json:"Hidden"`
		Enabled            bool            `json:"Enabled"`
		Readonly           bool            `json:"Readonly"`
		FItems             [][]interface{} `json:"F_Items"`
		SelectedValueArray []string        `json:"SelectedValueArray"`
	} `json:"p1_ddlShi"`
	P1DdlXian struct {
		Hidden             bool            `json:"Hidden"`
		Enabled            bool            `json:"Enabled"`
		FItems             [][]interface{} `json:"F_Items"`
		SelectedValueArray []string        `json:"SelectedValueArray"`
	} `json:"p1_ddlXian"`
	P1DdlJieDao struct {
		Hidden             bool            `json:"Hidden"`
		FItems             [][]interface{} `json:"F_Items"`
		SelectedValueArray []string        `json:"SelectedValueArray"`
	} `json:"p1_ddlJieDao"`
	P1XiangXDZ struct {
		Hidden bool   `json:"Hidden"`
		Label  string `json:"Label"`
		Text   string `json:"Text"`
	} `json:"p1_XiangXDZ"`
	P1ShiFZJ struct {
		Required      bool            `json:"Required"`
		Hidden        bool            `json:"Hidden"`
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue interface{}     `json:"SelectedValue"`
	} `json:"p1_ShiFZJ"`
	P1GaoZDFXLJS struct {
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue interface{}     `json:"SelectedValue"`
	} `json:"p1_GaoZDFXLJS"`
	P1QueZHZJC struct {
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue interface{}     `json:"SelectedValue"`
	} `json:"p1_QueZHZJC"`
	P1DangRGL struct {
		SelectedValue string          `json:"SelectedValue"`
		FItems        [][]interface{} `json:"F_Items"`
	} `json:"p1_DangRGL"`
	P1GeLSM struct {
		Hidden           bool `json:"Hidden"`
		IFrameAttributes struct {
		} `json:"IFrameAttributes"`
	} `json:"p1_GeLSM"`
	P1GeLFS struct {
		Required      bool            `json:"Required"`
		Hidden        bool            `json:"Hidden"`
		FItems        [][]interface{} `json:"F_Items"`
		SelectedValue interface{}     `json:"SelectedValue"`
	} `json:"p1_GeLFS"`
	P1GeLDZ struct {
		Hidden bool `json:"Hidden"`
	} `json:"p1_GeLDZ"`
	P1Ctl01BtnReturn struct {
		OnClientClick string `json:"OnClientClick"`
	} `json:"p1_ctl01_btnReturn"`
	P1 struct {
		Title            string `json:"Title"`
		IFrameAttributes struct {
		} `json:"IFrameAttributes"`
	} `json:"p1"`
}
type PostReportInfo struct {
	CoreInfo          *CoreInfo
	ViewState         string
	Fstatedaytemplate *F_STATE_DAY
	Time              time.Time
	XingCM            string
}

func (p *PostReportInfo) GetFstate() string {
	C := p.CoreInfo
	fstate := p.Fstatedaytemplate
	// fstate['p1_BaoSRQ']['Text'] = BaoSRQ
	// fstate['p1_P_GuoNei_ShiFSH']['SelectedValue'] = ShiFSH
	// fstate['p1_P_GuoNei_JinXXQ']['SelectedValueArray'][0] = JinXXQ
	// fstate['p1_P_GuoNei_ShiFZX']['SelectedValue'] = ShiFZX
	// fstate['p1_P_GuoNei_XiaoQu']['SelectedValue'] = XiaoQu
	// fstate['p1_ddlSheng']['F_Items'] = [[ddlSheng, ddlSheng, 1, '', '']]
	// fstate['p1_ddlSheng']['SelectedValueArray'] = [ddlSheng]
	// fstate['p1_ddlShi']['F_Items'] = [[ddlShi, ddlShi, 1, '', '']]
	// fstate['p1_ddlShi']['SelectedValueArray'] = [ddlShi]
	// fstate['p1_ddlXian']['F_Items'] = [[ddlXian, ddlXian, 1, '', '']]
	// fstate['p1_ddlXian']['SelectedValueArray'] = [ddlXian]
	// fstate['p1_ddlJieDao']['F_Items'] = [[ddlJieDao, ddlJieDao, 1, '', '']]
	// fstate['p1_ddlJieDao']['SelectedValueArray'] = [ddlJieDao]
	// fstate['p1_XiangXDZ']['Text'] = XiangXDZ
	// fstate['p1_ShiFZJ']['SelectedValue'] = ShiFZJ
	// fstate['p1_P_GuoNei_pImages_HFimgXingCM']['Text'] = XingCM
	fstate.P1BaoSRQ.Text = p.Time.Format("2006-01-02")
	fstate.P1PGuoNeiShiFSH.SelectedValue = C.ShiFSH
	fstate.P1PGuoNeiJinXXQ.SelectedValueArray[0] = C.JinXXQ
	fstate.P1PGuoNeiShiFZX.SelectedValue = C.ShiFZX
	fstate.P1PGuoNeiXiaoQu.SelectedValue = C.XiaoQu
	fstate.P1DdlSheng.FItems = [][]interface{}{{C.ddlSheng, C.ddlSheng, 1, "", ""}}
	fstate.P1DdlSheng.SelectedValueArray = []string{C.ddlSheng}
	fstate.P1DdlShi.FItems = [][]interface{}{{C.ddlShi, C.ddlShi, 1, "", ""}}
	fstate.P1DdlShi.SelectedValueArray = []string{C.ddlShi}
	fstate.P1DdlXian.FItems = [][]interface{}{{C.ddlXian, C.ddlXian, 1, "", ""}}
	fstate.P1DdlXian.SelectedValueArray = []string{C.ddlXian}
	fstate.P1DdlJieDao.FItems = [][]interface{}{{C.ddlJieDao, C.ddlJieDao, 1, "", ""}}
	fstate.P1DdlJieDao.SelectedValueArray = []string{C.ddlJieDao}
	fstate.P1XiangXDZ.Text = C.XiangXDZ
	fstate.P1ShiFZJ.SelectedValue = C.ShiFZJ
	fstate.P1PGuoNeiPImagesHFimgXingCM.Text = p.XingCM
	str, _ := json.Marshal(fstate)
	// log.Println(string(str))
	encodestate := base64.StdEncoding.EncodeToString(str)
	t := int(len(encodestate) / 2)
	// + "F_STATE" +
	encodestate = encodestate[:t] + encodestate[t:]
	return encodestate
}

func (s *PostReportInfo) GetKvs() map[string]string {
	BaoSRQ := s.Time.Format("2006-01-02")
	ViewState := s.ViewState
	F_STATE := s.GetFstate()
	ShiFSH := s.CoreInfo.ShiFSH
	JinXXQ := s.CoreInfo.JinXXQ
	ShiFZX := s.CoreInfo.ShiFZX
	XiaoQu := s.CoreInfo.XiaoQu
	XingCM := s.XingCM
	ddlSheng := s.CoreInfo.ddlSheng
	ddlShi := s.CoreInfo.ddlShi
	ddlXian := s.CoreInfo.ddlXian
	ddlJieDao := s.CoreInfo.ddlJieDao
	XiangXDZ := s.CoreInfo.XiangXDZ
	ShiFZJ := s.CoreInfo.ShiFZJ
	kvs := make(map[string]string)
	kvs["__EVENTTARGET"] = "p1$ctl01$btnSubmit"
	kvs["__EVENTARGUMENT"] = ""
	kvs["__VIEWSTATE"] = ViewState
	kvs["__VIEWSTATEGENERATOR"] = "7AD7E509"
	kvs["p1$ChengNuo"] = "p1_ChengNuo"
	kvs["p1$BaoSRQ"] = BaoSRQ
	kvs["p1$CengFWSS"] = "否"
	kvs["p1$DangQSTZK"] = "良好"
	kvs["p1$TiWen"] = ""
	kvs["p1$GuoNei"] = "国内"
	kvs["p1$P_GuoNei$ShiFSH"] = ShiFSH
	kvs["p1$P_GuoNei$JinXXQ"] = JinXXQ
	kvs["p1$P_GuoNei$ShiFZX"] = ShiFZX
	kvs["p1$P_GuoNei$XiaoQu"] = XiaoQu
	kvs["p1$P_GuoNei$pImages$HFimgXingCM"] = XingCM
	kvs["p1$JiuYe_ShouJHM"] = ""
	kvs["p1$JiuYe_Email"] = ""
	kvs["p1$JiuYe_Wechat"] = ""
	kvs["p1$QiuZZT"] = ""
	kvs["p1$JiuYKN"] = ""
	kvs["p1$JiuYSJ"] = ""
	kvs["p1$ddlGuoJia$Value"] = "-1"
	kvs["p1$ddlGuoJia"] = "选择国家"
	kvs["p1$ddlSheng$Value"] = ddlSheng
	kvs["p1$ddlSheng"] = ddlSheng
	kvs["p1$ddlShi$Value"] = ddlShi
	kvs["p1$ddlShi"] = ddlShi
	kvs["p1$ddlXian$Value"] = ddlXian
	kvs["p1$ddlXian"] = ddlXian
	kvs["p1$ddlJieDao$Value"] = ddlJieDao
	kvs["p1$ddlJieDao"] = ddlJieDao
	kvs["p1$XiangXDZ"] = XiangXDZ
	kvs["p1$ShiFZJ"] = ShiFZJ
	kvs["p1$GaoZDFXLJS"] = "无"
	kvs["p1$QueZHZJC"] = "否"
	kvs["p1$DangRGL"] = "否"
	kvs["p1$GeLDZ"] = ""
	kvs["p1$Address2"] = ""
	kvs["F_TARGET"] = "p1_ctl01_btnSubmit"
	kvs["p1_pnlDangSZS_Collapsed"] = "false"
	kvs["p1_P_GuoNei_pImages_Collapsed"] = "false"
	kvs["p1_P_GuoNei_Collapsed"] = "false"
	kvs["p1_GeLSM_Collapsed"] = "false"
	kvs["p1_Collapsed"] = "false"
	kvs["F_STATE"] = F_STATE
	return kvs
}

type CoreInfo struct {
	ShiFSH, JinXXQ, ShiFZX, XiaoQu, ddlSheng, ddlShi, ddlXian, ddlJieDao, XiangXDZ, ShiFZJ string
}

type SelfReportClient struct {
	Client *RedirectClient
}

func (s *SelfReportClient) Report(t time.Time) (*CoreInfo, error) {
	body := s.GetReportInfo(t)
	if s.cheakReportInfo(body, t) {
		return s.ParseReportInfo(body), nil
	} else {
		coreInfo, err := s.Report(t.Add(-time.Hour * 24))
		if err != nil {
			return nil, fmt.Errorf("%s 填报失败 %v", t.Add(-time.Hour*24).Format("2006-01-02"), err)
		}
		err = s.PostReport(coreInfo, t)
		return coreInfo, err
	}
}

func (s *SelfReportClient) GetReportInfo(t time.Time) []byte {
	year, mouth, day := cntime.GetYearMonthDay(t)
	url := fmt.Sprintf("https://selfreport.shu.edu.cn/ViewDayReport.aspx?day=%s-%s-%s", year, mouth, day)
	res, _ := s.Client.Get(url)
	body, _ := ioutil.ReadAll(res.Body)
	return body
}

func (s *SelfReportClient) cheakReportInfo(body []byte, t time.Time) bool {
	year := t.Year()
	mouth := t.Month()
	day := t.Day()

	if strings.Contains(string(body), "无指定日期的信息") {
		log.Printf("%v-%v-%v  无填报信息\n", year, mouth, day)
		return false
	} else {
		log.Printf("%v-%v-%v  有填报信息\n", year, mouth, day)
		return true
	}
}

func (s *SelfReportClient) ParseReportInfo(body []byte) *CoreInfo {
	ShiFSH, JinXXQ, ShiFZX, XiaoQu, ddlSheng, ddlShi, ddlXian, ddlJieDao, XiangXDZ, ShiFZJ := "在上海（校内）",
		"宝山", "是", "宝山", "上海", "上海市", "宝山区", "大场镇", "上海大学", "是"

	var re = regexp.MustCompile(`(?mU)var f.{1,2}_state=(.*var.*);`)
	Matchs := re.FindAllStringSubmatch(string(body), -1)
	var infoList = make([]string, 0, 30)
	for _, match := range Matchs {
		infoList = append(infoList, match[1])
	}
	getfromSelectedValueArray := func(str string) (string, bool) {
		var re = regexp.MustCompile(`(?mU)"SelectedValueArray":\["(.*)"\]`)
		strs := re.FindStringSubmatch(str)
		if len(strs) != 0 {
			return strs[1], true
		} else {
			return "", false
		}
	}
	getfromSelectedValue := func(str string) (string, bool) {
		var re = regexp.MustCompile(`(?mU)"SelectedValue":"(.*)"`)
		strs := re.FindStringSubmatch(str)
		if len(strs) != 0 {
			return strs[1], true
		} else {
			return "", false
		}
	}
	getfromText := func(str string) (string, bool) {
		var re = regexp.MustCompile(`(?mU)"Text":"(.*)"`)
		strs := re.FindStringSubmatch(str)
		if len(strs) != 0 {
			return strs[1], true
		} else {
			return "", false
		}

	}

	for _, info := range infoList {
		if strings.Contains(info, "BaoSRQ") {
			BaoSRQ, err := getfromText(info)
			if err != false {
				log.Print("-BaoSRQ-")
				log.Print(BaoSRQ)
			}

		}
		if strings.Contains(info, "ShiFSH") {

			ShiFSH0, err := getfromText(info)
			if err != false {
				log.Print("-ShiFSH-")
				ShiFSH = ShiFSH0
				log.Print(ShiFSH)
			}

		}
		if strings.Contains(info, "JinXXQ") {
			if strings.Contains(ShiFSH, "不在上海") {
				log.Print("-JinXXQ-")
				JinXXQ = ""
				log.Print(JinXXQ)
			} else {
				JinXXQ1, err := getfromText(info)
				if err != false {
					log.Print("-JinXXQ-")
					JinXXQ = JinXXQ1
					log.Print(JinXXQ)
				}
			}

		}
		if strings.Contains(info, "ShiFZX") {

			ShiFZX1, err := getfromSelectedValue(info)
			if err != false {
				log.Print("-ShiFZX-")
				ShiFZX = ShiFZX1
				log.Print(ShiFZX)
			}

		}
		if strings.Contains(info, "XiaoQu") {
			if strings.Contains(ShiFSH, "不在上海") || ShiFZX == "否" {
				log.Print("-XiaoQu-")
				XiaoQu = ""
				log.Print(XiaoQu)
			} else {
				XiaoQu1, err := getfromText(info)
				if err != false {
					log.Print("-XiaoQu-")
					XiaoQu = XiaoQu1
					log.Print(XiaoQu)
				}
			}

		}
		if strings.Contains(info, "ddlSheng") {

			ddlSheng1, err := getfromSelectedValueArray(info)
			if err != false {
				log.Print("-ddlSheng-")
				ddlSheng = ddlSheng1
				log.Print(ddlSheng)
			}

		}
		if strings.Contains(info, "ddlShi") {

			ddlShi1, err := getfromSelectedValueArray(info)
			if err != false {
				log.Print("-ddlShi-")
				ddlShi = ddlShi1
				log.Print(ddlShi)
			}

		}
		if strings.Contains(info, "ddlXian") {

			ddlXian1, err := getfromSelectedValueArray(info)
			if err != false {
				log.Print("-ddlXian-")
				ddlXian = ddlXian1
				log.Print(ddlXian)
			}

		}
		if strings.Contains(info, "ddlJieDao") {

			if strings.Contains(ShiFSH, "不在上海") {
				log.Print("-ddlJieDao-")
				ddlJieDao = ""
				log.Print(ddlJieDao)
			} else {
				ddlJieDao1, err := getfromSelectedValueArray(info)
				if err != false {
					log.Print("-ddlJieDao-")
					ddlJieDao = ddlJieDao1
					log.Print(ddlJieDao)
				}
				if ddlJieDao == "-1" {
					log.Print("-ddlJieDao-")
					ddlJieDao = "大场镇"
					log.Print(ddlJieDao)
				}
			}
		}
		if strings.Contains(info, "XiangXDZ") {
			ddlXian1, err := getfromText(info)
			if err != false {
				log.Print("-XiangXDZ-")
				XiangXDZ = ddlXian1
				log.Print(XiangXDZ)
			}

		}
		if strings.Contains(info, "ShiFZJ") {

			ShiFZJ1, err := getfromSelectedValue(info)
			if err != false {
				log.Print("-ShiFZJ-")
				ShiFZJ = ShiFZJ1
				log.Print(ShiFZJ)
			}

		}
	}
	if strings.Contains(ShiFSH, "不在") {
		log.Print("-JinXXQ-")
		JinXXQ = ""
		log.Print(JinXXQ)
	}
	if strings.Contains(ShiFSH, "不在") || ShiFZX == "否" {
		log.Print("-XiaoQu-")
		XiaoQu = ""
		log.Print(XiaoQu)
	}
	log.Println("")
	coreInfo := &CoreInfo{ShiFSH: ShiFSH, JinXXQ: JinXXQ, ShiFZX: ShiFZX,
		XiaoQu: XiaoQu, ddlSheng: ddlSheng, ddlShi: ddlShi, ddlXian: ddlXian, ddlJieDao: ddlJieDao, XiangXDZ: XiangXDZ, ShiFZJ: ShiFZJ}
	return coreInfo
}
func (s *SelfReportClient) PostReport(coreInfo *CoreInfo, t time.Time) error {
	buf := new(bytes.Buffer)
	bw := multipart.NewWriter(buf)
	view_state := s.GetViewState(t)
	phone := s.GetPhoneNum()
	// log.Println(view_state)
	xingCM := s.GetXingCM(phone, view_state, t)
	pri := &PostReportInfo{CoreInfo: coreInfo, ViewState: view_state, Fstatedaytemplate: s.GetFstatedaytemplate(), Time: t, XingCM: xingCM}
	kvs := pri.GetKvs()
	Addmultipartkvs(bw, kvs)
	bw.Close()
	req, _ := NewRequest("POST", "https://selfreport.shu.edu.cn/DayReport.aspx", buf)
	req.Header.Add("Content-Type", bw.FormDataContentType())
	req.Header.Add("X-FineUI-Ajax", "true")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	res, err := s.Client.BanRedirectDo(req)
	if err != nil {
		return fmt.Errorf("提交数据失败 %v", err)
	}
	body, _ := ioutil.ReadAll(res.Body)
	// log.Println(string(body))
	if strings.Contains(string(body), "提交成功") || strings.Contains(string(body), "历史信息不能修改") || strings.Contains(string(body), "现在还没到晚报时间") || strings.Contains(string(body), "只能填报当天或补填以前的信息") {
		log.Printf("每日一报Post%s成功\n", t.Format("2006-01-02"))
		return nil
	} else {
		log.Printf("每日一报Post%s失败\n", t.Format("2006-01-02"))
		log.Println(string(body))
		return errors.New(fmt.Sprintf("每日一报Post%s失败\n", t.Format("2006-01-02")))
	}

}
func (s *SelfReportClient) GetPhoneNum() string {
	PhoneNum := "111111111"
	req, _ := s.Client.Get("https://selfreport.shu.edu.cn/PersonInfo.aspx")
	str, _ := ioutil.ReadAll(req.Body)
	var re = regexp.MustCompile(`(?mU)var f.._state={"Text":"([\s\S]*)"}`)
	match := re.FindStringSubmatch(string(str))
	PhoneNum = match[1]
	return PhoneNum
}

func GetFstateimagetemplate() *F_STATE_IMAGE {
	fstateimage := new(F_STATE_IMAGE)
	fstateFile, err := ioutil.ReadFile("Resources/fstate_image.json")
	if err != nil {
		log.Println("读取模板失败", err)
	}
	json.Unmarshal(fstateFile, fstateimage)
	return fstateimage
}

func (s *SelfReportClient) GetFstatedaytemplate() *F_STATE_DAY {
	fstateday := new(F_STATE_DAY)
	fstateFile, err := ioutil.ReadFile("Resources/fstate_day.json")
	if err != nil {
		log.Println("读取模板失败", err)
	}
	json.Unmarshal(fstateFile, fstateday)
	return fstateday
}

func (s *SelfReportClient) GetViewState(t time.Time) string {
	year, month, day := cntime.GetYearMonthDay(t)
	url := fmt.Sprintf("https://selfreport.shu.edu.cn/DayReport.aspx?day=%s-%s-%s", year, month, day)
	rep, _ := s.Client.Get(url)
	str, _ := ioutil.ReadAll(rep.Body)
	retstr := s.parseViewState(string(str))
	return retstr
}
func (s *SelfReportClient) parseViewState(str string) string {
	var re = regexp.MustCompile(`(?mU)id="__VIEWSTATE" value="([\S\s]*)" \/>`)
	matchlist := re.FindAllStringSubmatch(str, -1)
	if len(matchlist) == 0 {
		log.Printf("没有找到ViewState")
		return ""
	}
	var retstr string
	for _, match := range matchlist {
		retstr = match[1]
	}
	return retstr
}

func (s *SelfReportClient) GetXingCM(PhoneNum string, ViewState string, t time.Time) string {
	var xingCM = "Xtw9cLnyqDla1iypPwrPtQ=="
	var re = regexp.MustCompile(`(?m)(.{22}==)|[\W]([A-Za-z0-9/]{11}=)`)
	res, _ := s.Client.Get("https://selfreport.shu.edu.cn/DayReport.aspx")
	body, _ := ioutil.ReadAll(res.Body)
	code := re.FindStringSubmatch(string(body))
	if code == nil {
		log.Println("没有找到行程码，正在上传行程码")
		err := creatXingCMimage(PhoneNum, t)
		if err != nil {
			log.Println(err, "创建图片失败")
			return xingCM
		}
		imageinfo := &Imageinfo{ViewState: ViewState, Time: t, Fstateimagetemplate: GetFstateimagetemplate()}
		contType, Reader, err := prepareImgmultipart(PhoneNum, imageinfo)
		if err != nil {
			log.Println(err, "打开文件失败，使用默认行程码")
			return xingCM
		}
		req, _ := NewRequest("POST", "https://selfreport.shu.edu.cn/DayReport.aspx", Reader)
		req.Header.Add("Content-Type", contType)
		req.Header.Add("X-Requested-With", "XMLHttpRequest")
		req.Header.Add("X-FineUI-Ajax", "true")
		res, err := s.Client.BanRedirectDo(req)
		if err != nil {
			log.Println("图片上传失败，使用默认行程码")
			return xingCM
		}
		body, _ := ioutil.ReadAll(res.Body)
		code := re.FindStringSubmatch(string(body))
		if code == nil {
			log.Println("未找到图片，使用默认行程码")
			log.Println(string(body))
			return xingCM
		}
		return code[1]
	}
	xingCM = code[1]
	return xingCM
}

func prepareImgmultipart(PhoneNum string, imageinfo *Imageinfo) (string, io.Reader, error) {
	img, err := ioutil.ReadFile(fmt.Sprintf("Resources/%s.jpeg", PhoneNum))
	if err != nil {
		return "", nil, err
	}
	buf := new(bytes.Buffer)
	bw := multipart.NewWriter(buf)
	fw1, _ := bw.CreateFormField("__EVENTTARGET")
	fw1.Write([]byte("p1$P_GuoNei$pImages$fileXingCM"))
	fw2, _ := bw.CreateFormField("__VIEWSTATE")
	fw2.Write([]byte(imageinfo.ViewState))
	fw3, _ := bw.CreateFormField("F_STATE")
	fw3.Write([]byte(imageinfo.GetFstate()))
	log.Println(imageinfo.GetFstate())
	filew, _ := bw.CreateFormFile("p1$P_GuoNei$pImages$fileXingCM", "xingchengma.jpeg")
	filew.Write(img)
	bw.Close()

	return bw.FormDataContentType(), buf, nil
}

func creatXingCMimage(phoneNum string, t time.Time) error {
	img, err := gg.LoadJPG("Resources/xingcm.jpg")
	if err != nil {
		log.Println("读取行程码模板错误")
		return err
	}
	context := gg.NewContextForImage(img)

	if err := context.LoadFontFace("Resources/yahei.ttf", 36); err != nil { // 从本地加载字体文件
		log.Println("加载字体文件错误")
		return err
	}

	context.SetRGB255(39, 39, 39)
	str := phoneNum[:3] + "****" + phoneNum[len(phoneNum)-4:] + "的动态行程卡"
	w, h := context.MeasureString(str)
	context.DrawString(str, 414-w/2, 380+h/2)

	if err := context.LoadFontFace("Resources/yahei.ttf", 30); err != nil { // 从本地加载字体文件
		log.Println("加载字体文件错误")
		return err
	}
	context.SetRGB255(143, 142, 147)
	str = "更新于：" + t.Format("2006-01-02 03:04:05")
	context.DrawString(str, 414-w/2, 460+h/2)
	image := context.Image()
	p := fmt.Sprintf("Resources/%s.jpeg", phoneNum)
	f, err := os.OpenFile(p, os.O_SYNC|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	err = jpeg.Encode(f, image, &jpeg.Options{Quality: 80})
	if err != nil {
		return err
	}
	return nil
}
