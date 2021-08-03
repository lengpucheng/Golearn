package main

/*
cmd tcp 命令
telnet ip port
*/
func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
