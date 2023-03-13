package service

import (
	"adsec-wx-bot/common"
	"adsec-wx-bot/dao"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

var baseUrl = "https://api.github.com/search/repositories?q="
var baseParam = "&sort=updated"

func GitSearch(keyWords string) (gitMaps []dao.GitDAO, err error) {
	searchUrl := baseUrl + keyWords + baseParam
	resp, err := http.Get(searchUrl)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	responseBody, _ := io.ReadAll(resp.Body)
	git := &dao.Git{}
	err = json.Unmarshal(responseBody, git)
	if git.TotalCount != 0 {
		items := git.Items
		for _, item := range items {
			if int(time.Now().Sub(item.PushedAt).Hours()) <= 24 {
				isBlackKey := 0
				for _, key := range common.BlackKeyWords {
					if strings.Contains(item.Description, key) {
						isBlackKey++
					}
				}
				if isBlackKey != 0 {
					break
				}
				gitDao := dao.GitDAO{}
				gitDao.HTMLUrl = item.HtmlUrl
				gitDao.Name = item.Name
				gitDao.Description = item.Description
				gitDao.PushedAt = item.PushedAt
				gitMaps = append(gitMaps, gitDao)
			}
		}
	}
	return gitMaps, err
}
