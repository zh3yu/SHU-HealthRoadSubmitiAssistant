# SHU-HealthRoadSubmitiAssistant
上海大学健康之路提交助手，该项目参考 [@BlueFisher](https://github.com/BlueFisher/SHU-selfreport) 在该项目的基础上，用go做了重构，增加了申请第二天离校的功能。

## 每日一报
从当前日期开始查询是否有日报记录，若无记录则向前查询，直到有记录为止。用该记录填报之后的所有日期。

**！！！不支持进行虚假填报！！！**

位置发生改变后请及时手动申请每日一报。
## 离校申请

学校系统改版后。若满足三天两报的条件，可以申请一次第二天离校一次。该离校会被系统**自动**审核,不需要学校额外审核。有该申报后即可在一天内出门**一次**。

每日离校申请，可模拟疫情前的生活，自由出校。

## 使用方法
使用之前需要更改配置文件 [AccountList.yaml](Configuration/AccountList.yaml)
![image/配置文件](image/%E9%85%8D%E7%BD%AE%E6%96%87%E4%BB%B6.png)
若要使用测试用例进行测试，请配置[core.go](core/core.go)
![image/测试用例配置](image/%E6%B5%8B%E8%AF%95%E7%94%A8%E4%BE%8B%E9%85%8D%E7%BD%AE.png)
### win直接运行程序
右侧下载release，解压后运行。
**win自动化运行** 
没有采取[@BlueFisher](https://github.com/BlueFisher/SHU-selfreport)用github Action的方案。考虑到github服务器在国外，有连不上学校网的可能，需要配置openvpn。推荐使用win任务计划程序进行自动化申报。
![服务](image/%E8%87%AA%E5%8A%A8%E4%BB%BB%E5%8A%A1%E6%B5%81%E7%A8%8B1.png)
右侧选择创建任务，设置触发器和操作。
！[触发器](image/%E8%87%AA%E5%8A%A8%E4%BB%BB%E5%8A%A1%E6%B5%81%E7%A8%8B2.png)
！[操作](image/%E8%87%AA%E5%8A%A8%E4%BB%BB%E5%8A%A1%E6%B5%81%E7%A8%8B3.png)

**windows代码编译**
请确保计算机已经安装go运行环境。使用build.bat即可编译程序。 
### 在linux上运行
**docker** (推荐使用)

将项目down到服务器后,将项目打包成REPOSITORY。
```
docker build -t SHUHR:latest . 
docker run it SHUHR:latest
```
~~**linux代码编译**~~
~~没有测试，应该也可以。~~

**linux自动化运行**
可用 crontab 进行计划任务。
```
crontab -e 
```
![时间格式](image/%E6%97%B6%E9%97%B4%E6%A0%BC%E5%BC%8F.png)

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
