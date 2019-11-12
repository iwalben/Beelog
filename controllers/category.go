package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this * CategoryController)Get()  {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login",302)
		return
	}

	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsCategory"] = true
	this.TplName = "category.html"
	this.Data["Categories"],_ = models.GetAllCategories()
}

func (this *CategoryController)Post()  {
	inputs := this.Input()
	_ = models.AddCategory(inputs.Get("title"))
	this.Redirect("/category",301)
}