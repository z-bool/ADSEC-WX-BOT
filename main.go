package main

import (
	botService "adsec-wx-bot/service/bot"
	"flag"
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

func init() {
	fmt.Println(`
 ________  ________  ________  _______   ________                ________  ________  _________   
|\   __  \|\   ___ \|\   ____\|\  ___ \ |\   ____\              |\   __  \|\   __  \|\___   ___\ 
\ \  \|\  \ \  \_|\ \ \  \___|\ \   __/|\ \  \___|  ____________\ \  \|\ /\ \  \|\  \|___ \  \_| 
 \ \   __  \ \  \ \\ \ \_____  \ \  \_|/_\ \  \    |\____________\ \   __  \ \  \\\  \   \ \  \  
  \ \  \ \  \ \  \_\\ \|____|\  \ \  \_|\ \ \  \___\|____________|\ \  \|\  \ \  \\\  \   \ \  \ 
   \ \__\ \__\ \_______\____\_\  \ \_______\ \_______\             \ \_______\ \_______\   \ \__\
    \|__|\|__|\|_______|\_________\|_______|\|_______|              \|_______|\|_______|    \|__|
                       \|_________|                                                              
									欢迎使用ADSEC-WX-BOT，关注公众号：阿呆攻防，获取更多信息                                               
                                                                                                 

`)
}

func main() {
	var lineMode bool
	var startTime string
	flag.BoolVar(&lineMode, "lineMode", false, "是否以命令行模式开启二维码,默认以链接形式(default false)")
	flag.StringVar(&startTime, "startTime", "8:00", "定时任务时间设置，默认为8:00")
	flag.Parse()
	self, bot := botService.CreateBot(lineMode)
	if self != nil && bot != nil {
		// 群组推送实时订阅
		groups, _ := self.Groups()
		timezone, _ := time.LoadLocation("Asia/Shanghai")
		s := gocron.NewScheduler(timezone)
		s.Every(1).Day().At(startTime).Do(func() {
			botService.SendIntelligence(groups)
		})
		s.StartBlocking()
		bot.Block()
	}

}
