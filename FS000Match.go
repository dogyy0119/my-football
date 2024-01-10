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
)

func main() {

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
