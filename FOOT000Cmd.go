package main

import (
	"bufio"
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
	"strings"
	launch2 "tesou.io/platform/foot-parent/foot-core/launch"
	"tesou.io/platform/foot-parent/foot-spider/launch"
	"time"
)

func main() {
	go My_Spider_timer()
	go Spider_timer()

HEAD:
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	input = strings.ToLower(input)
	switch input {
	case "exit\n", "exit\r\n", "quit\n", "quit\r\n":
		break
	case "\n", "\r\n":
		goto HEAD
	case "init\n", "init\r\n":
		launch2.GenTable()
		//launch2.TruncateTable()
		goto HEAD
	case "spider\n", "spider\r\n":
		launch.Spider()
		goto HEAD
	case "new_spider\n", "new_spider\r\n":
		launch.My_Spider()
		goto HEAD
	case "analy\n", "analy\r\n":
		launch2.Analy(false)
		goto HEAD
	case "new_analy\n", "new_analy\r\n":
		//launch2.Analy(false)
		launch2.Analy_new(false)
		//
		//launch2.Analy_Near()
		goto HEAD
	default:
		goto HEAD
	}

}

func My_Spider_timer() {
	fmt.Println(" My_Spider_timer. ")
	ticker := time.NewTicker(15 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			launch.My_Spider()
			//launch2.Analy_new(false)
		}
	}
}

func Spider_timer() {
	fmt.Println(" Spider_timer. ")
	ticker := time.NewTicker(20 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			launch.Spider()

			//launch2.Analy(false)
		}
	}
}

func Analy_timer() {
	fmt.Println(" Analy_timer.")
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			launch2.Analy(false)
			launch2.Analy_new(false)
		}
	}
}

func sendmail(body string) {

	m := gomail.NewMessage()
	//发送人
	m.SetHeader("From", "499489735@qq.com")
	//接收人
	m.SetHeader("To", "499489735@qq.com")
	//抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
	//主题
	m.SetHeader("Subject", "小佩奇")
	//内容
	m.SetBody("text/html", body)
	//附件
	//m.Attach("./myIpPic.png")

	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer("smtp.qq.com", 587, "499489735@qq.com", "uwalxsdkwvjvbicd")

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("DialAndSend err %v:", err)
		panic(err)
	}
	fmt.Printf("send mail success\n")
}
