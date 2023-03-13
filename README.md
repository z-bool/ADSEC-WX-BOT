# ADSEC-WX-BOT 微信开源威胁情报机器人
**郑重声明：文中所涉及的技术、思路和工具仅供以安全为目的的学习交流使用，<u>任何人不得将其用于非法用途以及盈利等目的，否则后果自行承担</u>** 。
<p align="center"><a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-_red.svg"></a><a href="https://github.com/z-bool/ADSEC-WX-BOT"><img  src="https://goreportcard.com/badge/github.com/projectdiscovery/httpx"></a></p>

<p align="center"><a href="#install">模块打包</a> · <a href="#tall">使用说明</a> · <a href="#notice">注意事项</a> · <a href="#communicate">技术交流</a></p>

<div id="install"></div>

### 模块打包

```bash
go mod init adsec-wx-bot # 如果低版本不能执行的话可以删除go.mod文件重新init一下
go mod tidy # 加载依赖
go build . # 打包
```

如果没有二开或自行打包需要的，可以直接从release下载打包好的文件即可

<div id= "tall"></div>

### 使用说明

此机器人为开源简单版的威胁情报机器人，功能比较单一。开发的目的是为了获取每日的威胁情报信息用于跟进1day或者nday信息。默认设置的定时任务是每天早上8点，抓取的数据就是如:`2023.03.13 8:00 - 2023.03.14 8:00`这样也就不需要保存到数据库中，简化部署成本。更多功能其实是为了社群管理服务的，更多功能可以加入阿呆的`技术交流群`体验。

```text
adsec-wx-bot -h # 查看帮助信息
adsec-wx-bot -lineMode true # 开启命令行扫码登录,适用于linux
adsec-wx-bot # 默认使用链接扫码登录
adsec-wx-bot -lineMode true -startTime "12:00" # 命令行扫码并且设置12点的定时任务
```
<p>帮助信息:</p>
![help](https://github.com/z-bool/ADSEC-WX-BOT/blob/main/img/img.png)
<p>链接登录:</p>
![链接登录](https://github.com/z-bool/ADSEC-WX-BOT/blob/main/img/img1.png)
<p>命令行二维码:</p>
![二维码登录](https://github.com/z-bool/ADSEC-WX-BOT/blob/main/img/img3.png)
<p>结果展示:</p>
![结果展示](https://github.com/z-bool/ADSEC-WX-BOT/blob/main/img/img2.png)
<p>如果要增加监控的关键词，请修改`common/common.go`中的白名单关键词</p>
![关键词](https://github.com/z-bool/ADSEC-WX-BOT/blob/main/img/img4.png)
<div id="notice"></div>

### 注意事项

- 使用时请将要推送的群聊`保存到通讯录`
- 一次登录长期使用，会在运行目录下生成`cookie.json`，重视文件安全，防止被窃取该文件

<div id="communicate"></div>

### 技术交流

<img src="https://cdn.jsdelivr.net/gh/z-bool/images@master/img/qrcode_for_gh_c90beef1e2e7_258.jpg" alt="阿呆攻防公众号" style="zoom:100%;" />![image-20230116173105809](https://cdn.jsdelivr.net/gh/z-bool/images@master/img/image-20230116173105809.png)



微信群有过期时间限制，如果有技术交流、BUG解决、环境安装问题都可以于公众号/QQ群获取微信群信息。