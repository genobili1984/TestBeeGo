package routers

import (
	"github.com/genobili1984/TestBeeGo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
