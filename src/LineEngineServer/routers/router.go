package routers

import (
	"LineEngineServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.RESTRouter("/object", &controllers.ObjectController{})
	beego.RESTRouter("/openBook/:path", &controllers.OpenBookController{})
}
