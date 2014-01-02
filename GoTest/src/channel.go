package main

import (
	"fmt"
)

/**
注意，channel默认上是阻塞的，
也就是说，如果Channel满了，就阻塞写，
如果Channel空了，就阻塞读。
于是，我们就可以使用这种特性来同步我们的发送和接收端。
**/
func channelSample() {
	//默认buffer为1， make(chan string, 2)  //这样就变成了2
	channel := make(chan string)

	go func() { channel <- "hello" }()

	msg := <-channel
	fmt.Println(msg)
}

func main() {
	channelSample()
}
