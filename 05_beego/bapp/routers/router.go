package routers

import (
	"github.com/advanced-go-coding/05_beego/bapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
