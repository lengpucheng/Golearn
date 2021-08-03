package main

import "net"

type User struct {
	Name string
	Addr string
	Chan chan string
	conn net.Conn
}

// 创建用户
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name: "u#" + userAddr,
		Addr: userAddr,
		Chan: make(chan string),
		conn: conn,
	}
	// 启动消息监听协程
	go user.ListenMessage()

	return user
}

// 监听当前的chan 有消息发送给 user
func (t *User) ListenMessage() {
	for {
		// 从管道中读取
		msg := <-t.Chan
		// 写入消息
		t.conn.Write([]byte(msg + "\n"))
	}
}
