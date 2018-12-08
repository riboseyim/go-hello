package main

import (
	"log"

	"github.com/go-gomail/gomail"
)

const (
	Sender       = "zijingshanke@qq.com"
	SenderPass   = "vfmwxurvzqyqbdgi"
	SenderServer = "smtp.qq.com"
	SenderPort   = 587
	Recivers     = "yanrui@zhong-ying.com"
)

func TTankMail(to string, Subject string, Content string) (bool, error) {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", Sender /*"发件人地址"*/, "来自严睿的邮箱") // 发件人
	// 收件人

	m.SetHeader("To",
		m.FormatAddress(to, "收件人"))

	//抄送
	//m.SetHeader("Cc", m.FormatAddress("kurui_alarm@163.com", "收件人"))

	//	m.SetHeader("Bcc",m.FormatAddress("xxxx@gmail.com", "收件人")) / 暗送

	m.SetHeader("Subject", Subject) // 主题
	m.SetBody("text/html", Content) // 正文
	//m.Attach("我是附件")
	d := gomail.NewPlainDialer(SenderServer, SenderPort, Sender, SenderPass)
	// 发送邮件服务器、端口、发件人账号、发件人密码

	if err := d.DialAndSend(m); err != nil {
		log.Println("===SendMail() Failed...", err)
		return false, err
	}
	log.Println("===SendMail() Success..")
	return true, nil
}
