package main

import (
	"net"
	"time"

	"github.com/zzh20/timewheel"
)

var wheelHeartbeat = timewheel.New(time.Second*1, 30, func(data interface{}) {
	c := data.(net.Conn)
	c.Close()
})

func main() {
	ch := make(chan bool)
	wheelHeartbeat.Start()
	<-ch
}

// 连接建立完成后，连接对象加入时间轮中
//wheelHeartbeat.Add(conn)

// 收到客户端发送的心跳包后，连接对象加入时间轮中
//wheelHeartbeat.Add(conn)


