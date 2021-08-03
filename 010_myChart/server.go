package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	// IP
	Ip string
	// 端口
	Port int
	// message 管道
	MessageChan chan string
	// 用户表 储存用户的引用 非指针
	OlineUser map[string]*User
	// map读写锁
	lock sync.RWMutex
}

// 创建一个server实例 返回值为 server的指针  -- 类似于构造器
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:          ip,
		Port:        port,
		MessageChan: make(chan string),
		OlineUser:   make(map[string]*User),
	}
	return server
}

// 监听Message广播消息
func (t *Server) ListenMessage() {
	for {
		// 读取消息
		msg := <-t.MessageChan
		// 发送给全部的user
		t.lock.Lock()
		for _, u := range t.OlineUser {
			u.Chan <- msg
		}
		t.lock.Unlock()
	}
}

// 广播消息
func (t *Server) BroadMsg(user *User, msg string) {
	// 拼接消息
	sendMsg := "[" + user.Addr + "]" + user.Name + " : " + msg
	// 写入消息到管道
	t.MessageChan <- sendMsg
}

// 业务逻辑处理器
func (t *Server) Handler(conn net.Conn) {
	// 初始化用户
	user := NewUser(conn)
	// 用户加入在线列表
	t.lock.Lock()
	t.OlineUser[user.Name] = user
	t.lock.Unlock()
	// 广播用户上线消息
	fmt.Println("【日志】来自[" + user.Addr + "]的链接已建立")
	t.BroadMsg(user, "我已上线，可以和我聊天了！")
}

// 启动服务
func (t *Server) Start() {
	// socket listen 监听套接字
	// net.Listen可以初始化监听一个端口 参数为 协议（TCP/UDP） 和 地址（ip：port)
	// fmt.Sprintf 可以按照格式 格式化字符串
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", t.Ip, t.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
	}

	// close 关闭
	defer listen.Close()

	// 启动监听message
	go t.ListenMessage()

	// 一直监听全部链接
	for {
		// accept 接收 conn为建立链接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen accept err:", err)
			continue
		}
		// do handler 启动协程去异步处理
		go t.Handler(conn)
	}

}
