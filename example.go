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
