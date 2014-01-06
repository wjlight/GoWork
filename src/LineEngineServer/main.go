package main

import (
	"LineEngineServer/models"
	_ "LineEngineServer/routers"
	"github.com/astaxie/beego"
)

func init() {
	// rService = rpcService()
	models.RpcInit("127.0.0.1", "9090")
}

//		Objects

//	URL					HTTP Verb				Functionality
//	/object				POST					Creating Objects
//	/object/<objectId>	GET						Retrieving Objects
//	/object/<objectId>	PUT						Updating Objects
//	/object				GET						Queries
//	/object/<objectId>	DELETE					Deleting Objects

func main() {
	beego.Run()
}
