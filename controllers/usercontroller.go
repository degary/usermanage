package controllers

import (
	"fmt"
	"github.com/degary/usermanage/forms"
	"github.com/degary/usermanage/services"
)

type UserController struct {
	BaseController
}

//abc@123   $2a$10$jwNOOay.UE9kHmn5mdfmDeIWFT7JUQMC99RKWv3bMe1gA5SybIoV2

func (c *UserController)List()  {
	users := services.GetUsers()
	c.Data["users"] = users
	c.TplName="usercontroller/list.html"

}

func (c *UserController)Prepare()  {
	c.BaseController.Prepare()
	c.Data["navKey"] = "user"

}

func (c *UserController)Delete()  {
	if getInt64, err := c.GetInt64("id");err == nil {
		services.DeleteUser(getInt64)
	}
	c.Redirect("/user/list",302)
}

func (c *UserController)Add()  {
	updateuserform := &forms.UpdateUserForm{}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(updateuserform);err != nil {
			c.Redirect("/user/list",302)
			return
		}
		if updateuserform.Password1 != updateuserform.Password2 {
			c.Data["errorPasswd"] = "两次输入密码不一致" //errorPasswd
			fmt.Println(updateuserform.Name)
			c.Data["name"] = updateuserform.Name
			fmt.Println(updateuserform.Address)
			c.Data["address"] = updateuserform.Address
			c.Redirect("/user/add",302)
			return
		}
		services.AddUser(updateuserform.Name,updateuserform.Password1,updateuserform.Address)
		c.Redirect("/user/list",302)
	}
	c.TplName = "usercontroller/add.html"
}