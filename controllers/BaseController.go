package controllers

import (
	"github.com/astaxie/beego"
	"github.com/degary/usermanage/services"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController)Prepare()  {
	c.Data["currentUser"] = nil
	c.Data["navKey"] = ""

	id :=c.GetSession("user")
	if id == nil {
		//未登录
		c.Redirect("/auth/login/",302)
		return
	}
	if pk,ok := id.(int64);ok {
		//如果user不为nil,则返回当前user
		if 	user := services.GetUserById(pk);user != nil {
		c.Data["currentUser"] = user
		}
	}

	if c.Data["currentUser"] == nil {
		c.DestroySession()
		c.Redirect("/auth/login",302)
		return
	}




}
