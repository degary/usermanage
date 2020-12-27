package controllers

type ArticleController struct {
	BaseController
}

func (Article *ArticleController)List()  {
	//获取名字为"user"的session
	user :=Article.GetSession("user")
	if user == nil {
		Article.Redirect("/auth/login/",302)
		return
	}
}

func (c *ArticleController)Prepare()  {
	c.BaseController.Prepare()
	c.Data["navKey"] = "article"
}
