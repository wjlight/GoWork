package main

import (
	"fmt"
)

type Server interface {
	Name() string
}

type IpcServer struct {
	Server
}

func (i *IpcServer) Name() string {
	fmt.Println("wawa")
	return "wawa"
}

func main() {
	ipc := &IpcServer{}
	ipc.Name()
	fmt.Println("hello world")
}
