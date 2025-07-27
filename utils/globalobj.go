package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"zinx/ziface"
)

type GlobalObj struct {
	// Server

	TcpServer ziface.IServer //  全局Server对象
	Host      string         // IP
	TcpPort   int            // 端口号
	Name      string         //服务器名称

	// Zinx

	Version          string // Zinx版本
	MaxConn          int    //允许最大连接数
	MaxPacketSize    uint32 // 数据包最大值
	WorkerPoolSize   uint32 // 当前业务工作Woker池的Goroutine数量
	MaxWorkerTaskLen uint32 // Zinx框架允许最大的Woker数量

}

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		fmt.Println("read conf/zinx.json fail")
		panic(err)
	}
	// 解析Json数据
	err = json.Unmarshal(data, &g)
	if err != nil {
		return
	}
}

var GlobalObject *GlobalObj

func init() {
	// 默认文件
	GlobalObject = &GlobalObj{
		Host:             "0.0.0.0",
		TcpPort:          8888,
		Name:             "zinx",
		Version:          "V0.4",
		MaxConn:          1000,
		MaxPacketSize:    4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
	}

	// 读取用户定义参数
	//GlobalObject.Reload()
}
