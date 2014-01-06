package models

import (
	"LineEngineServer/models/service/demo"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"os"
)

var rService demo.WorkBooks
var rTransport *thrift.TSocket

func RpcInit(rpcIp string, rpcPort string) {
	rService, rTransport = rpcService(rpcIp, rpcPort)
}

//"127.0.0.1", "9090"
func rpcService(rpcIp string, rpcPort string) (demo.WorkBooks, *thrift.TSocket) {
	fmt.Printf("start rpc ...\n")
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(rpcIp, rpcPort))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}
	userTransport := transportFactory.GetTransport(transport)
	service := demo.NewWorkBooksClientFactory(userTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:9090", " ", err)
		os.Exit(1)
	}

	// defer transport.Close()

	// filePath := "record03.xls"
	// back, err := service.Open(filePath)
	fmt.Printf("end rpc ...\n")
	return service, transport
}

func GetService() demo.WorkBooks {
	return rService
}
