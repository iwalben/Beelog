package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["IsHome"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["Topics"],_ = models.GetAllTopics(true)
	this.TplName = "home.html"
}