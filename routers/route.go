package routers

import (
	"github.com/astaxie/beego"
	"github.com/degary/usermanage/controllers"
)

func init()  {
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.ArticleController{})
	beego.AutoRouter(&controllers.AuthController{})
}
