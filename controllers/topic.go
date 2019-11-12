package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController)Get()  {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login",302)
		return
	}

	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = "topic.html"
	this.Data["IsTopic"] = true

	topics , err := models.GetAllTopics(false)
	if err != nil{
		beego.Error(err)
	}else {
		this.Data["Topics"]=topics
	}
}

func (this *TopicController)Post()  {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login",302)
		return
	}
	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")

	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title,content)
	}else {
		err = models.ModifyTopic(tid,title,content)
	}

	if err != nil{
		beego.Error(err)
	}
	this.Redirect("/topic",302)
}


//智能路由规则

func (this *TopicController)Add()  {
	this.TplName = "topic_add.html"
}

func (this *TopicController)Del()  {
	//0代表第一个参数
	id := this.Ctx.Input.Param("0")
	_ = models.DelTopic(id)
	this.Redirect("/topic",302)
}

func (this *TopicController)Show()  {
	//0代表第一个参数
	tid := this.Ctx.Input.Param("0")
	topic , err := models.GetTopic(tid)
	if err != nil {
		this.Redirect("/topic",302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["id"] = tid
	this.TplName = "topic_view.html"
}

func (this * TopicController)Modify()  {
	tid := this.Ctx.Input.Param("0")
	topic , err := models.GetTopic(tid)
	if err != nil {
		this.Redirect("/topic",302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["id"] = tid
	this.TplName = "topic_modify.html"
}


