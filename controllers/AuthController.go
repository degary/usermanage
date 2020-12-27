package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/degary/usermanage/services"
	"github.com/degary/usermanage/forms"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController)Login()  {
	loginform := &forms.LoginForm{}
	var errMsg string
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(loginform);err == nil{
			logs.Info(loginform)
			if user := services.Auth(loginform);user != nil {
				logs.Info("用户:%s登录成功",user.Name)
				//存储session 名字为 "user"
				c.SetSession("user",user.ID)
				c.Redirect("/article/list",302)
				return
			}else {
				errMsg = "用户名密码错误"
			}
		}
	}
	c.Data["form"] = nil
	c.Data["error"] = errMsg
	c.TplName="authcontroller/login.tpl"
}

func (c *AuthController)Logout() {
	c.DestroySession()
	c.Redirect("/auth/login/",302)
}
