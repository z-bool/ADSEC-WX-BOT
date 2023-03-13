package service

import (
	"adsec-wx-bot/common"
	spiderService "adsec-wx-bot/service/spider"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/skip2/go-qrcode"
	"time"
)

// 打印二维码到控制台
func consoleQrCode(uuid string) {
	q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
	fmt.Println(q.ToString(true))
}

// 机器人登录回调
func loginBot(bot *openwechat.Bot) *openwechat.Bot {
	reloadStorage := openwechat.NewJsonFileHotReloadStorage("cookie.json")
	defer reloadStorage.Close()

	if err := bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
		fmt.Println(err)
	}
	return bot
}

// 创建一个机器人会话
func CreateBot(lineMode bool) (*openwechat.Self, *openwechat.Bot) {
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式，上面登录不上的可以尝试切换这种模式
	if lineMode {
		bot.UUIDCallback = consoleQrCode
	} else {
		bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	}
	bot = loginBot(bot)

	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	return self, bot
}

// 发送威胁情报的任务
func SendIntelligence(groups openwechat.Groups) {
	for _, group := range groups {
		for _, key := range common.GitSearchKeyWords {
			gitMaps, _ := spiderService.GitSearch(key)
			if gitMaps != nil {
				group.SendText(GitNoticePush(gitMaps, key))
			}
			time.Sleep(5 * time.Second)
		}
	}
}
