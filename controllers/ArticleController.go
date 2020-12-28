package controllers

import (
	"fmt"
	"github.com/degary/usermanage/forms"
	"github.com/degary/usermanage/services"
)

type ArticleController struct {
	BaseController
}

func (c *ArticleController) List() {
	articles := services.GetArticles()
	c.Data["articles"] = articles
	c.TplName = "articlecontroller/list.tpl"
}

func (c *ArticleController) Delete() {
	if getInt64, err := c.GetInt64("id"); err != nil {
		c.Data["err"] = err
		c.TplName = "articlecontroller/list.tpl"
		return
	} else {
		err = services.DeleteArticleById(getInt64)
		c.Data["err"] = err
		c.TplName = "articlecontroller/list.tpl"
	}
}

func (c *ArticleController) Update() {
	var articleForm = &forms.UpdateArticleForm{}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(articleForm); err != nil {
			c.Redirect("/article/list", 302)
			return
		}
		if err := services.UpdateArticle(articleForm); err == nil {
			c.Redirect("/article/list", 302)
		} else {
			c.Data["articleForm"] = articleForm
			c.TplName = "acticlecontroller/list"
		}
	}
	id, _ := c.GetInt64("id")
	fmt.Println(id)
	article := services.GetArticleById(id)
	c.Data["article"] = &article
	c.TplName = "articlecontroller/update.html"
}

func (c *ArticleController) Prepare() {
	c.BaseController.Prepare()
	c.Data["navKey"] = "article"
}
