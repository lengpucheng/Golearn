package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIP   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(ip string, port int) *Client {
	// 创建链接
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// 创建客户端
	client := &Client{ip, port, "unName", conn, 1}
	return client
}

// menu菜单
func (the *Client) menu() bool {
	var flag int
	fmt.Println("1. 广播模式")
	fmt.Println("2. 在线列表")
	fmt.Println("3. 私聊模式")
	fmt.Println("4. 修改名称")
	fmt.Println("0. 退出程序")

	_, err := fmt.Scanln(&flag)
	if err != nil {
		return false
	}

	if flag >= 0 && flag <= 4 {
		the.flag = flag
		return true
	} else {
		LogSys(">>>>>>>请输入正确的项目代码<<<<<,")
		return false
	}
}

// Run 运行
func (the *Client) Run() {
	the.Send("--help")
	LogSys("使用 --help 可以获取帮助")
	var msg string
	for {
		_, err := fmt.Scanln(&msg)
		if err != nil {
			return
		}
		if len(msg) > 1 {
			the.Send(msg)
		}
	}
}

// RunMENU 带菜单运行
func (the *Client) RunMENU() {
	for the.flag != 0 {
		for !the.menu() {

		}
		switch the.flag {
		case 1:
			Log(">>>>>广播模式<<<<<")
			break
		case 2:
			Log(">>>>>在线列表<<<<<")
			break
		case 3:
			Log(">>>>>私聊模式<<<<<")
			break
		case 4:
			Log(">>>>>修改名称<<<<<")
			break
		}
	}
}

// Send 发送消息
func (the *Client) Send(msg string) {
	buffer := []byte(msg + "\n")
	_, err := the.conn.Write(buffer)
	if err != nil {
		LogSys("【ERR】消息发送失败")
		return
	}
}

// Show 显示消息
func (the *Client) Show() {
	// 一旦conn有数据，就将其拷贝到stdout中，否则一直阻塞
	_, err := io.Copy(os.Stdout, the.conn)
	if err != nil {
		LogSys("显示消息错误", err)
		return
	}
}
