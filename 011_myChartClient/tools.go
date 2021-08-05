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

// Log 日志
func Log(msg ...interface{}) {
	fmt.Println(msg)
}

// LogSys 系统日志
func LogSys(msg ...interface{}) {
	Log(fmt.Sprintf("【SYS】%s:", GetNowTime()), msg)
}
