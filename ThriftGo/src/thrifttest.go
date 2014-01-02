package main

import (
	"Hello"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"os"
	"time"
)

func main() {
	startTime := currentTimeMillis()

	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "9090"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}
	userTransport := transportFactory.GetTransport(transport)
	client := Hello.NewHelloClientFactory(userTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:9090", " ", err)
		os.Exit(1)
	}

	defer transport.Close()

	back, err := client.HelloString("wawa")
	fmt.Println("Response from server :" + back)

	endTime := currentTimeMillis()

	fmt.Println("Program exit. time->", endTime, startTime, (endTime - startTime))

}

// 转换成毫秒
func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
