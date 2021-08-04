package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name    string
	Addr    string
	Channel chan string
	conn    net.Conn
	server  *Server
}

// NewUser 构造
func NewUser(conn net.Conn, server *Server) *User {
	addr := conn.RemoteAddr().String()
	deferName := "user-" + GetPortForAddr(addr)
	user := &User{deferName, addr, make(chan string), conn, server}

	// 启动消息监听
	go user.ListenMsg()

	return user
}

// ListenMsg 从chan中读取信息回传
func (the *User) ListenMsg() {
	for {
		// 从chan中读取数据
		msg, ok := <-the.Channel
		if !ok {
			// Chan 被关闭后就退出
			return
		}
		// 写入发送
		_, err := the.conn.Write([]byte(msg))
		if err != nil {
			LogUser(the, "发送失败", err)
		}
	}
}

// GetMsgBefore 获取用户消息前缀
func (the *User) GetMsgBefore() string {
	return fmt.Sprintf("[%s]%s(%s):", GetNowTime(), the.Name, the.Addr)
}

// Oline 用户上线
func (the *User) Oline() {
	// 写入用户表
	the.server.lock.Lock()
	the.server.Users[the.Name] = the
	the.server.lock.Unlock()
	// 写日志并广播
	the.server.SendMsg(the, LogMsg("System", the, "我上线了"))
}

// OffOline 下线
func (the *User) OffOline() {
	// 去除用户
	the.server.lock.Lock()
	delete(the.server.Users, the.Name)
	the.server.lock.Unlock()
	//广播
	the.server.SendMsg(the, LogMsg("System", the, "断开链接"))
}

// Handler 处理
func (the *User) Handler(msg string) {
	if msg == "--help" {
		the.MsgCallBack(Help())
	} else if msg == "--list" {
		list := "-----在线列表-----"
		for _, user := range the.server.Users {
			list += "\n" + fmt.Sprintf("[%s](%s)", user.Name, user.Addr)
		}
		list += "\n------------------\n"
		the.MsgCallBack(list)
	} else if len(msg) > 9 && msg[:9] == "--rename " {
		name := msg[9:]
		// 更新表
		the.server.lock.Lock()
		_, ok := the.server.Users[name]
		if ok {
			the.MsgCallBack(LogMsg("System", the, "名称已被占用", name))
		} else {
			delete(the.server.Users, the.Name)
			the.Name = name
			the.server.Users[name] = the
			// 日志
			the.MsgCallBack(LogMsg("System", the, "名称修改成功"))
		}
		the.server.lock.Unlock()
	} else if len(msg) > 5 && msg[:5] == "--to " {
		// 获取发送的用户名
		index := strings.Index(msg, ">")
		if index == -1 || len(msg) < index+1 {
			the.MsgCallBack("命令参数错误，请使用[--help]查询用法\n" + Help())
			return
		}
		un := msg[5:index]
		message := msg[index+1:]
		// 获取用户
		user, ok := the.server.Users[un]
		if !ok {
			the.MsgCallBack(fmt.Sprintf("用户[%s]不存在,请使用[--list]获取用户列表\n", un))
			return
		}
		// 发送消息
		user.MsgCallBack(the.GetMsgBefore() + message)
	} else {
		the.server.SendMsg(the, msg+"\n")
	}
}

// MsgCallBack 消息回调
func (the *User) MsgCallBack(msg string) {
	_, err := the.conn.Write([]byte(msg))
	if err != nil {
		LogUser(the, "推送消息失败", err)
		return
	}
}
