package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type DeleteController struct {
	beego.Controller
}

func (this * DeleteController)Get()  {
	id := this.Ctx.Input.Params()[":id"]
	_ = models.DelCategory(id)
	this.Redirect("/category",301)
}