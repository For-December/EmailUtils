package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func dealLine(f func(param ...string)) {

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	// 使用 bufio.NewScanner 创建一个 Scanner 来逐行读取文件内容
	scanner := bufio.NewScanner(file)

	// 循环读取文件内容
	for scanner.Scan() {
		// 使用 strings.Split 函数按制表符分隔行数据
		line := scanner.Text()
		fields := strings.Split(line, "\t")

		// 处理每个字段
		f(fields...)
	}

	// 检查是否有错误发生
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件失败:", err)
	}
}
