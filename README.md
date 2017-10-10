# 时间轮算法

go语言实现的时间轮算法，用于定时任务的处理。

常用场景：tcp连接保持所用的心跳包机制

tcp连接检测常用心跳包机制，客户端定时向服务器发送一个标示心跳的空包，服务器有几种方式处理：
  1. 每个连接创建一个定时器处理，超时时间内收到客户端的心跳包，认为客户端连接有效，启动下一个定时器。
  2. 创建一个hashmap储存所有连接对象，对象保存下一次心跳超时的时间，收到心跳包后更新对应连接的超时时间，启动一个定时器，定时扫描（遍历）map，断开所有已经超时的连接。
  
问题：服务器端连接数量很大时，会占用过多的系统资源及影响服务器的响应时间。

timewheel算法，只需要启动一个定时器，所有定时任务加入相应的时间槽内。

<img src="http://img.my.csdn.net/uploads/201209/29/1348926970_9123.png" alt="">

定时轮的工作原理可以类比于时钟，如上图箭头（指针）按某一个方向按固定频率轮动，每一次跳动称为一个 tick。

这样可以看出定时轮由个3个重要的属性参数，ticksPerWheel（一轮的tick数），tickDuration（一个tick的持续时间）

以及 timeUnit（时间单位），例如 当ticksPerWheel=60，tickDuration=1，timeUnit=秒，这就和现实中的始终的秒针走动完全类似了。


这里给出一种简单的实现方式，指针按 tickDuration 的设置进行固定频率的转动，其中的必要约定如下：

    新加入的对象总是保存在当前指针转动方向上一个位置
    相等的对象仅存在于一个 slot 中
    指针转动到当前位置对应的 slot 中保存的对象就意味着 timeout 了
    

# 使用

go get -u github.com/zzh20/timewheel

# 例子

```go

package main

import (
  "net"
  "log"
  
  "github.com/zzh20/timewheel"
)

// 定义心跳包，设置心跳超时时间，处理函数
var wheelHeartbeat = timewheel.New(time.Second*1, 30, func(data interface{}) {
	c := data.(net.Conn)
	log.Printf("timeout close conn:%v", c)
	c.Close()
})

func main() {

  // 启动心跳包检查
  wheelHeartbeat.Start()
  
}

// 客户端连接成功 
func SessionConnected() {
    wheelHeartbeat.Add(conn)
}

// 客户端连接断开
func SessionClosed() {
    wheelHeartbeat.Remove(conn))
}

// 处理客户端的心跳包
func HeartbeatHandler() {
  wheelHeartbeat.Add(conn)
}

```
