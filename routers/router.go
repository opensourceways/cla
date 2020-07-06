package routers

import (
	"github.com/opensourceways/cla/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
