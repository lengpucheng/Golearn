package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

const (
	// 允许的最大时间
	MaxTime = time.Second * 60
)

type Server struct {
	Ip   string
	Port int
	// 消息总线
	MsgBus chan string
	// 用户表
	Users map[string]*User
	// 读写锁
	lock sync.RWMutex
}

// NewServer 构造
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:     ip,
		Port:   port,
		MsgBus: make(chan string),
		Users:  make(map[string]*User),
		lock:   sync.RWMutex{},
	}

	return server
}

// Start 启动
func (the *Server) Start() {
	// 监听端口
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", the.Ip, the.Port))
	if err != nil {
		fmt.Println("start listen err:", err)
		return
	}

	// defer 关闭监听
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			fmt.Println("close listen err：", err)
		}
	}(listen)

	// 启动消息监听
	go the.ListenMsgBus()
	TitleInfo()

	// 一直监听
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept conn err:", err)
			continue
		}
		// 处理器
		go the.Handle(conn)
	}
}

// Handle 处理器
func (the *Server) Handle(conn net.Conn) {
	// 1. 创建用户
	user := NewUser(conn, the)
	// 2. 用户上线
	user.Oline()
	isLife := make(chan bool)
	// 3. 启动消息监听
	go the.ListenMsg(user, isLife)
	// 4. 超时检查
	for {
		select {
		// 当isLife出现 会进入分子
		case <-isLife:
			// 该分支为空 会重新执行select after被刷新
		// 相当于一个管道 一段时间后会有数据进入
		case <-time.After(MaxTime):
			user.MsgCallBack(fmt.Sprintf("超时%d秒未响应，已断开链接", MaxTime/time.Second))
			delete(the.Users, user.Name)
			// 强制退出
			close(user.Channel)
			err := conn.Close()
			if err != nil {
				LogUser(user, "关闭失败", err)
				return
			}
			// 退出当前用户的处理器
			return
		}
	}
}

// SendMsg 广播消息
func (the *Server) SendMsg(user *User, msg string) {
	sendMsg := fmt.Sprintf("%s:%s", user.GetMsgBefore(), msg)
	the.MsgBus <- sendMsg
}

// ListenMsg 监听消息
func (the *Server) ListenMsg(user *User, left chan bool) {
	for {
		buffer := make([]byte, 4096)
		size, err := user.conn.Read(buffer)
		if size == 0 {
			user.OffOline()
			return
		}
		if err != nil && err != io.EOF {
			LogUser(user, " send msg err ", err)
			continue
		}
		// 存活判断
		left <- true
		// 处理消息
		user.Handler(string(buffer[:size-1]))
	}
}

// ListenMsgBus 监听消息总线
func (the *Server) ListenMsgBus() {
	for {
		// 取出消息
		msg := <-the.MsgBus
		// 广播
		the.lock.Lock()
		for _, user := range the.Users {
			user.Channel <- msg
		}
		the.lock.Unlock()
	}

}
