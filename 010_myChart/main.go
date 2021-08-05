package main

/*
cmd tcp 命令
telnet ip port
*/
func main() {
	Log("myChart启动")
	NewServer("172.28.80.1", 8888).Start()
}
