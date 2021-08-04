package main

/*
cmd tcp 命令
telnet ip port
*/
func main() {
	Log("myChart启动")
	NewServer("127.0.0.1", 8888).Start()
}
