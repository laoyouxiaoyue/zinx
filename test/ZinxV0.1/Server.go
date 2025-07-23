package main

import "zinx/znet"

// Zinx框架实现服务端测试

func main() {
	s := znet.NewServer("[zinx V0.1]")
	s.Start()
}
