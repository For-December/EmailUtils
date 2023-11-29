package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/smtp"
	"strings"
	"time"
)

var smtpServer string
var smtpPort string
var senderEmail string
var senderName string
var auth smtp.Auth
var mailType string

func init() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("请您添加配置文件（将 config.yaml.demo 重命名为 config.yaml 并补充相应字段）")
		return // 自动退出
	}

	smtpServer = viper.GetString("email.smtpServer")
	smtpPort = viper.GetString("email.smtpPort")
	senderEmail = viper.GetString("email.senderEmail")
	senderName = "计算机学院学习部"
	senderPassword := viper.GetString("email.senderPassword")
	println(senderPassword)
	auth = smtp.PlainAuth("",
		senderEmail,
		senderPassword,
		smtpServer)
	mailType = "Content-Type: text/plain; charset=UTF-8"

}

const FormatNotice = `亲爱的 %s 同学：
    欢迎你参加本次联合组织的四/六级模拟考。
您的考场为: %s 号
座位号为: %s 号
考试时间: 2023年12月2日（本周六）15：00开考。

【温馨提示】
1. 14：30开始入场，14：40试音，请在试音前进入考场
2. 本次模拟考采用教室电脑外放听力（不需要带耳机）
3. 文具自备（包括2B铅笔，黑色签字笔等等）
4. 考试结束后试卷（含听力）及答案将在QQ群公布,请关注群消息

请大家严格遵守正式四六级考试的相关规定，提前做好准备，确保在规定时间内到达考场。
预祝大家考试顺利！谢谢合作！
		——计算机学院学生会学习部
`

func SendNoticeEmail(param ...string) {
	if len(param) != 4 {
		log.Fatalln("参数个数不合法！")
	}
	stuName := param[0]
	addrNum := param[1]
	seatNum := param[2]
	email := param[3]
	body := fmt.Sprintf(FormatNotice, stuName, addrNum, seatNum)
	subject := "2023年12月2日（本周六）15：00四六级模拟通知"
	SendEmail(email, subject, body)
	log.Printf("%s => %s %s 已完成\n", senderName, email, stuName)

	time.Sleep(5 * time.Second)

}
func SendEmail(toEmail string, subject string, body string) {
	s := fmt.Sprintf("To:%s\r\n"+
		"From:%s <%s>\r\n"+
		"Subject:%s\r\n"+
		"%s\r\n\r\n%s",
		toEmail, senderName, senderEmail, subject, mailType, body)
	msg := []byte(s)

	err := smtp.SendMail(
		smtpServer+":"+smtpPort,
		auth, senderEmail,
		strings.Split(toEmail, ","),
		msg)
	if err != nil {
		log.Fatalln(err)
	}

}
