package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	pojo1 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/module/match/service"
	serv1 "tesou.io/platform/foot-parent/foot-core/module/odds/service"
)

func main() {
	servi1 := new(serv1.EuroTrackService)
	temp1 := new(pojo1.EuroTrack)
	temp1.MatchId = "2508883"
	temp1.CompId = 616
	temp1.OddDate = "2024-01-09 04:37:00"
	temp1.Num = 0

	exist, b := servi1.Exist(temp1)

	fmt.Println(exist)
	fmt.Println(b)

	servi := new(service.BFJinService)
	temp := new(pojo.BFJin)
	temp.ScheduleID = 2457953
	temp.HomeTeam = "侯城U21"
	temp.GuestTeam = "高云地利U21"
	temp.MatchTimeStr = "2023-09-26 02:00:00"

	jinSaveList := make([]interface{}, 0)
	jinSaveList = append(jinSaveList, temp)

	exist1, b1 := servi.Exist(temp)

	fmt.Println(exist1)
	fmt.Println(b1)
	if b1 {
		fmt.Println("存在")
	} else {
		jinSaveList = append(jinSaveList, temp)
		servi.SaveList(jinSaveList)
	}

	matchHisService := new(service.MatchHisService)
	temp2 := new(pojo.MatchHis)

	matchHisService.Exist(temp2)
	return
	resp, err := http.Get("http://m.titan007.com/analy/Analysis/2483939.htm")
	if err != nil {
		fmt.Println("http get error")
	}
	defer resp.Body.Close() // 需要关闭响应体

	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("htmlData error")
	}
	//fmt.Println(string(htmlData))
	// 使用 goquery 库

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlData))
	if err != nil {
		fmt.Println("fmt error")
	}

	// 查找指定标签元素

	doc.Find("script").Each(func(index int, item *goquery.Selection) {

		scriptContent := item.Text()
		if strings.Contains(scriptContent, "var jsonData") {
			// 找到含有 "var jsonData" 的 script 标签
			// 在这里处理您需要的逻辑

			jsonRegex := regexp.MustCompile(`var jsonData\s*=\s*(.+?);`)
			match := jsonRegex.FindStringSubmatch(scriptContent)

			//fmt.Println("Found script with var jsonData:", match)

			var jsonData map[string]interface{}

			if len(match) == 2 {
				//fmt.Println("Found script with var jsonData  match 0:", match[0])
				fmt.Println("Found script with var jsonData  match 1:", match[1])

				jsonDataStr := match[1]
				if err := json.Unmarshal([]byte(jsonDataStr), &jsonData); err != nil {
					panic(err)
				}
				if err := ioutil.WriteFile("D:/output.txt", []byte(jsonDataStr), 0644); err != nil {
					panic(err)
				}

			}
			//fmt.Println(jsonData)

		}

	})
}
