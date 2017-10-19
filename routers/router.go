package routers

import (
	"github.com/astaxie/beego"
	"xlsx-reading/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/upload", &controllers.MainController{}, "*:Upload")
}
