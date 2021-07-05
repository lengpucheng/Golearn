package main

// 导入必须使用 否则会报错
// 若在导入前加 _ 表示匿名导入  既可以不使用导入（调用 init方法）
// 可以给包取别名
// 若在导入前加 . 表示可以直接使用当前包内的方法  不需要限定名 尽量不用.
import (
	_ "Golearn/003_函数/liba"
	lib "Golearn/003_函数/libb"
)

func main() {
	lib.LibaInc()
}
