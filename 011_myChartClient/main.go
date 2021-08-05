package main

import (
	"flag"
	"fmt"
)

var (
	ip   string
	port int
)

func main() {
	// 命令行解析
	flag.Parse()

	if ip == "127.0.0.1" && port == 8888 {
		Log("请输入服务器IP:")
		_, err := fmt.Scanln(&ip)
		if err != nil {
			LogSys("输入错误")
			return
		}
		Log("请输入端口号:")
		_, err = fmt.Scanln(&port)
		if err != nil {
			LogSys("输入错误")
			return
		}
	}

	client := NewClient(ip, port)
	if client == nil {
		LogSys(">>>>>>>>>>>链接失败<<<<<<<<<")
		return
	}
	LogSys(">>>>>>>服务器链接成功<<<<<<<")
	// 消息显示
	go client.Show()
	// 客户端业务
	client.Run()
}

// init 解析命令行参数
func init() {
	flag.StringVar(&ip, "ip", "127.0.0.1", "设置服务器IP")
	flag.IntVar(&port, "port", 8888, "设置服务器端口")
}
