package main

import (
	"fmt"
	"strings"
	"time"
)

// GetPortForAddr 获取进程号
func GetPortForAddr(addr string) string {
	return addr[strings.Index(addr, ":")+1:]
}

// GetNowTime 获取当前时间 YYYY-MM-DD HH:mm:ss
func GetNowTime() string {
	return time.Now().String()[:19]
}

// Log 记录日志
func Log(msg ...interface{}) string {
	return LogUser(nil, msg)
}

// LogUser 记录日志
func LogUser(user *User, msg ...interface{}) string {
	return LogMsg("log", user, msg)
}

// LogMsg 消息日志
func LogMsg(mod string, user *User, msg ...interface{}) string {
	log := fmt.Sprintf("【%s】[%s]:", mod, GetNowTime())
	if user != nil {
		log = fmt.Sprintf("%s{%s}(%s)", log, user.Name, user.Addr)
	}
	log = fmt.Sprintln(log, msg)
	fmt.Println(log)
	return log
}

// Help 帮助菜单
func Help() (help string) {
	help = "--list"
	help += "\n  查看主线列表 用法【--list】"
	help += "\n--rename"
	help += "\n  修改用户名 用法【--rename [新名称]】"
	help += "\n--to"
	help += "\n  向指定用户发送消息 用法【--to [用户名]>[消息内容]】\n"
	return
}
