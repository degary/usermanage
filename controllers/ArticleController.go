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
		c.Redirect("/article/list/", 302)
		return
	} else {
		err = services.DeleteArticleById(getInt64)
		c.Data["err"] = err
		c.Redirect("/article/list/", 302)
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
	article := services.GetArticleById(id)
	c.Data["article"] = &article
	c.TplName = "articlecontroller/update.html"
}

func (c *ArticleController) Add() {
	if c.Ctx.Input.IsPost() {
		article := forms.UpdateArticleForm{}
		err2 := c.ParseForm(&article)
		if err2 != nil {
			c.Data["errMsg"] = err2
			c.TplName = "articlecontroller/add.html"
			return
		}
		err := services.AddArticle(article.Name, article.Author, article.Price)
		if err != nil {
			errMsg := fmt.Sprintf(err.Error())
			c.Data["errMsg"] = errMsg
		} else {
			c.Redirect("/article/list", 302)
		}
	}
	c.TplName = "articlecontroller/add.html"
}

func (c *ArticleController) Prepare() {
	c.BaseController.Prepare()
	c.Data["navKey"] = "article"
}
