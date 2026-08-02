# Time Wheel Algorithm

A time wheel algorithm implemented in Go for handling scheduled tasks.

Common scenario: TCP connection heartbeat mechanism.

For TCP connection detection, a common approach is using heartbeat packets. The client sends an empty packet indicating a heartbeat at regular intervals, and the server handles it in several ways:
  1. Create a timer for each connection. If a heartbeat packet is received within the timeout period, the client connection is considered valid, and the next timer is started.
  2. Create a hashmap to store all connection objects. Each object records the next heartbeat timeout time. Upon receiving a heartbeat packet, update the corresponding connection's timeout time. Start a timer to scan (iterate through) the map periodically, and disconnect all connections that have timed out.

Problem: When the number of server-side connections is very large, this approach consumes excessive system resources and affects server response time.

The time wheel algorithm only requires starting a single timer. All scheduled tasks are added to the appropriate time slot.

<img src="http://img.my.csdn.net/uploads/201209/29/1348926970_9123.png" alt="">

The working principle of the time wheel can be analogized to a clock. As shown in the figure, the arrow (pointer) rotates in a certain direction at a fixed frequency. Each rotation is called a tick.

This indicates that the time wheel has three important parameters: ticksPerWheel (number of ticks per full rotation), tickDuration (duration of a single tick), and timeUnit (time unit). For example, when ticksPerWheel=60, tickDuration=1, and timeUnit=seconds, this is analogous to the second hand of a clock moving.

Here is a simple implementation approach, where the pointer rotates at a fixed frequency as configured by tickDuration. The following necessary conventions are applied:

    Newly added objects are always stored in the slot in the direction of the pointer's rotation.
    Equal objects exist only in one slot.
    When the pointer reaches the slot corresponding to the current position, the objects stored in it are considered to have timed out.

# Usage

go get -u github.com/zzh20/timewheel

# Example

```go

package main

import (
  "net"
  "log"
  
  "github.com/zzh20/timewheel"
)

// Define heartbeat, set heartbeat timeout, processing function
var wheelHeartbeat = timewheel.New(time.Second*1, 30, func(data interface{}) {
	c := data.(net.Conn)
	log.Printf("timeout close conn:%v", c)
	c.Close()
})

func main() {

  // Start heartbeat checking
  wheelHeartbeat.Start()
  
}

// Client connection established 
func SessionConnected() {
    wheelHeartbeat.Add(conn)
}

// Client connection disconnected
func SessionClosed() {
    wheelHeartbeat.Remove(conn))
}

// Process client heartbeat packet
func HeartbeatHandler() {
  wheelHeartbeat.Add(conn)
}

```
</arg_value></tool_call>
