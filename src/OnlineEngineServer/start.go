package main

import (
	"encoding/json"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"service/demo"
)

func check(err error) {
	if err != nil {
		// panic(err)
		// log.Fatal("err:", err.Error())
		// os.Exit(1)
		fmt.Printf("error" + err.Error())
	} else {
		// log.Fatal("ok:")
		fmt.Printf("ok\n")
	}
}

// func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		defer func() {
// 			if e, ok := recover().(error); ok {
// 				http.Error(w, e.Error(), http.StatusInternalServerError)
// 				//可以自定义的错误显示
// 				//logging
// 				log.Println("WARN: panic in %v - %v", fn, e)
// 				log.Println(string(debug.Stack()))
// 			}
// 		}()
// 		fn(w, r)
// 	}
// }
var rService demo.WorkBooks

func init() {
	rService = rpcService()
}

func openHandler(rw http.ResponseWriter, req *http.Request) {
	// filePath := "record03.xls"
	filePath := req.FormValue("path")
	back, err := rService.Open(filePath)
	check(err)
	if back != nil {
		fmt.Printf("name: %s\n", back.Name)
	} else {
		fmt.Printf("rpc error")
	}
	s := toJson(back)
	rw.Header().Set("Content-Type", "application/json")
	io.WriteString(rw, s)
	// io.WriteString(rw, "<html><h1>"+s+"</h1></html>")
}

func toJson(entry interface{}) string {
	b, err := json.Marshal(entry)
	if err != nil {
		return ""
	}
	return string(b)
}

func rpcService() demo.WorkBooks {
	fmt.Printf("start rpc ...\n")
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "9090"))
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
	return service
}

func main() {
	http.HandleFunc("/openBook", openHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
