package service

import (
	"adsec-wx-bot/dao"
	"strconv"
)

func GitNoticePush(gitMaps []dao.GitDAO, keyword string) string {
	header := "亲爱的白帽子们，现在是【阿呆】为您播报今日威胁情报：\r\n"
	body := "【关键词】: #" + keyword + "\r\n"
	for index, item := range gitMaps {
		body += "------------项目" + strconv.Itoa(index+1) + "------------\r\n" +
			"【项目名称】" + item.Name + "\r\n" +
			"【项目地址】 " + item.HTMLUrl + "\r\n" +
			"【项目描述】 " + item.Description + "\r\n" +
			"【更新时间】 " + item.PushedAt.String() + "\r\n"
	}
	footer := "今日推送已完成，祝各位师傅漏洞挖到手软😄"
	return header + body + footer
}
