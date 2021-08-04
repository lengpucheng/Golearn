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
