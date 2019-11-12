package main

import (
	"beeblog/controllers"
	"beeblog/models"
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init()  {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	//force 是否每次重新建表
	//verbose 是否打印建表的相关信息
	orm.RunSyncdb("default",false,true)
	beego.Router("/",&controllers.MainController{})
	beego.Router("/category",&controllers.CategoryController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Router("/topic",&controllers.TopicController{})
	beego.Router("/delete/:id([0-9]+)",&controllers.DeleteController{})
	//TopicController添加智能路由
	beego.AutoRouter(&controllers.TopicController{})
	beego.Run()
}



