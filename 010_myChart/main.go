package main

/*
cmd tcp 命令
telnet ip port
*/
func main() {
	Log("myChart启动")
	NewServer("172.23.160.1", 8888).Start()
}
