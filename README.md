# SHU-HealthRoadSubmitiAssistant
上海大学健康之路提交助手，该项目参考 [@BlueFisher](https://github.com/BlueFisher/SHU-selfreport) 在该项目的基础上，用go做了重构，增加了申请第二天离校的功能。

## 每日一报
从当前日期开始查询是否有日报记录，若无记录则向前查询，直到有记录为止。用该记录填报之后的所有日期。

**！！！不支持进行虚假填报！！！**

位置发生改变后请及时手动申请每日一报。
## 离校申请

学校系统改版后。若满足三天两报的条件，可以申请一次第二天离校一次。该离校会被系统**自动**审核,不需要学校额外审核。有该申报后即可在一天内出门**一次**。

每日离校申请，可模拟疫情前的生活，自由出校。

##使用方法
TODO

## 免责声明
本项目仅作为免费的网络研究使用，
不得利用本程序以任何方式直接或者间接的从事违反中国法律、国际公约以及社会公德的行为，
**！！！不支持进行虚假填报！！！**

## 依赖
	[github.com/PuerkitoBio/goquery v1.8.0](github.com/PuerkitoBio/goquery)
	[github.com/fogleman/gg v1.3.0](github.com/fogleman/gg)
	[github.com/json-iterator/go v1.1.12](github.com/json-iterator/go)
	[gopkg.in/yaml.v2 v2.4.0](gopkg.in/yaml.v2)
  
  go结构体生成使用
  - **YAML** [@zhwt](https://github.com/Zhwt) [yaml-to-go](https://zhwt.github.io/yaml-to-go/)
  - **JSON** [@zhwt](https://github.com/Zhwt) [json-to-go](https://mholt.github.io/json-to-go/)
  
  其他在线工具使用
  - 正则表达式 [@firasdib](https://github.com/firasdib/Regex101) [Regex101](https://regex101.com/)
  - base64解码 @二环同学 [base64.us](https://base64.us/)
