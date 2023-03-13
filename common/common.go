package common

// 抓取数据时的黑名单，拒绝任意关键词的输出
var BlackKeyWords = [...]string{"反中共", "中共", "宪法", "大陆", "修宪", "中国共产党", "共党", "红色", "习近平", "江泽民", "胡锦涛"}

// 抓取数据时域名黑名单，拒绝携带的黑名单关键词
var BlackDomains = [...]string{"*.edu.cn", "*.gov.cn"}

// Git抓取数据时的白名单关键词
var GitSearchKeyWords = [...]string{"cve-2022", "cve-2023", "exploit", "GetShell", "RCE", "命令执行", "remote command execute", "代码执行", "免杀", "远控", "cobalt strike", "sql injection", "红队", "redteam"}
