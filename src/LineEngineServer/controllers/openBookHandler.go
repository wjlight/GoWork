package controllers

import (
	"LineEngineServer/models"
	"fmt"
	"github.com/astaxie/beego"
)

// type ResponseInfo struct {
// }

type OpenBookController struct {
	beego.Controller
}

func (this *OpenBookController) Get() {
	fmt.Printf("openBookController,,,,,,,")
	path := this.Ctx.Input.Params[":path"]
	rService := models.GetService()
	back, err := rService.Open(path)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = back
	}
	this.ServeJson()
}
