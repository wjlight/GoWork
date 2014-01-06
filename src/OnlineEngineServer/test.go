package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"os"
	"service/demo"
)

func main() {
	filePath := "record03.xls"
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "9090"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}
	userTransport := transportFactory.GetTransport(transport)
	client := demo.NewWorkBooksClientFactory(userTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:9090", " ", err)
		os.Exit(1)
	}

	defer transport.Close()

	back, err := client.Open(filePath)
	if err != nil {
		fmt.Println("Response from server err:", err.Error())
	} else {
		fmt.Println("Response from server sheet count:", back.SheetCount)
		// fmt.Println("Response from server sheet count:", back.SheetCount)
		fmt.Println("Response from server sheet count:", back.Name)
	}
}

// func test(a interface{}) {
// 	switch x := a.(type) {
// 	case demo.WorkBook:
// 		fmt.Println("x is WorkBook")
// 	case float64:
// 		fmt.Println("x is float64")
// 	}
// }
