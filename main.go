package main

import "log"

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	dealLine(SendNoticeEmail)
}

//SendEmail("1921567337@qq.com", "好好好", "你好呀~")
