package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"gopkg.in/gomail.v2"

	//"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

func main() {
	test2("test2")
	test1("test1")
}

func test2(body string) {
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

func test1(body string) {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = "liuhang <499489735@qq.com>"
	// 设置接收方的邮箱
	e.To = []string{"499489735@qq.com"}
	//设置主题
	e.Subject = "这是主题"
	//设置文件发送的内容
	e.HTML = []byte(body)
	//这块是设置附件
	e.AttachFile("./test.txt")
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "499489735@qq.com", "uwalxsdkwvjvbicd", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
}
