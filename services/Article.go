package services

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/degary/usermanage/forms"
	"github.com/degary/usermanage/modules"
)

func GetArticles() []*modules.Article {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&modules.Article{})
	var articles []*modules.Article
	queryset.All(&articles)
	return articles
}

func GetArticleById(pk int64) *modules.Article {
	ormer := orm.NewOrm()
	article := &modules.Article{Id: pk}
	if err := ormer.Read(article); err == nil {
		return article
	}
	return nil
}

func GetArticleByName(name string) *modules.Article {
	ormer := orm.NewOrm()
	article := &modules.Article{Name: name}
	if err := ormer.Read(&article); err == nil {
		return article
	}
	return nil
}

func DeleteArticleById(pk int64) error {
	ormer := orm.NewOrm()
	_, err := ormer.Delete(&modules.Article{Id: pk})
	return err
}

func UpdateArticle(form *forms.UpdateArticleForm) error {
	ormer := orm.NewOrm()
	if form.Id == 0 {
		return fmt.Errorf("%s", "id不能为空")
	}
	article := &modules.Article{
		Id:     form.Id,
		Name:   form.Name,
		Author: form.Author,
		Price:  form.Price,
	}
	if err := ormer.Read(&modules.Article{Id: form.Id}); err == nil {
		ormer.Update(article)
		return nil
	} else {
		return err
	}
}
