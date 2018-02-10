package routers

import (
	"github.com/advanced-go-coding/05_beego/bapi/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
