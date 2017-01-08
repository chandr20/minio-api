package routers

import (
	"minio_api_server/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/:id",&controllers.MainController{} )
}
