package main

import (
	"fmt"
	"log/slog"
	"net"
	"time"
)

// 模拟客户端测试

func main() {
	// 1 连接
	conn, err := net.Dial("tcp", "127.0.0.1:80")
	if err != nil {
		slog.Error(fmt.Sprint("conn err:", err))
		panic(err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			slog.Error(fmt.Sprint("conn err:", err))
		}
	}(conn)

	for {
		_, err := conn.Write([]byte("hello"))
		if err != nil {
			slog.Error(fmt.Sprint("write err:", err))
			return
		}
		slog.Info(fmt.Sprintf("write message: %s", "hello"))
		buf := make([]byte, 512)
		readLen, err := conn.Read(buf)
		if err != nil {
			slog.Error(fmt.Sprint("read err:", err))
			return
		}
		slog.Info(fmt.Sprintf("read message: %s", buf[:readLen]))

		time.Sleep(1 * time.Second)
	}

}
