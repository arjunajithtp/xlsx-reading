package routers

import (
	"fileUpload/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/upload", &controllers.MainController{}, "*:Upload")
}
