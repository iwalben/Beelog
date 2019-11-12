package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"time"
)

const(
	_DB_NAME = "data/beelogUFO.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id int64
	Title string
	Created time.Time `orm:"index"`
	Views int64 `orm:"index"`
	TopicTime time.Time `orm:"index"`
	TopicCount int64
	TopicLastUserId int64
}

type Topic struct {
	Id int64
	Uid int64
	Title string
	Content string `orm:"size(5000)"`
	Attachment string
	Created time.Time `orm:"index"`
	Updated time.Time `orm:"index"`
	Views int64 `orm:"index"`
	Author string
	ReplyTime time.Time `orm:"index"`
	Replycont int64
	ReplyLastUserId int64
}

func RegisterDB()  {
	if !com.IsExist(_DB_NAME) {
		_ = os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		_,_ = os.Create(_DB_NAME)
	}

	//注册模型
	orm.RegisterModel(new(Category),new(Topic))
	//注册驱动
	_ = orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	//注册默认数据库
	_ = orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddCategory(name string) error  {
	//创建orm
	o := orm.NewOrm()
	//创建Category
	cate := &Category{Title:name,Created:time.Now(),TopicTime:time.Now()}
	//查询
	qs := o.QueryTable("category")
	//筛选
	err := qs.Filter("title",name).One(cate)
	if err == nil{
		return err
	}
	//插入
	_ ,err = o.Insert(cate)
	if err != nil{
		return err
	}
	return nil
}

func GetAllCategories() ([]*Category,error) {
	o := orm.NewOrm()
	cates := make([]*Category ,0)
	qs := o.QueryTable("category")
	_,err := qs.All(&cates)
	return cates,err
}

func DelCategory(id string) error  {
	cid ,err := strconv.ParseInt(id,10,64)
	if err != nil{
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id:cid}
	_,err = o.Delete(cate)
	return err
}

func AddTopic(title , content string) error {
	o := orm.NewOrm()

	topic := &Topic{
		Title:title,
		Content:content,
		Created:time.Now(),
		Updated:time.Now(),
		ReplyTime:time.Now(),
	}
	_,err := o.Insert(topic)
	return err
}

func GetAllTopics(isDesc bool)([]*Topic ,error)  {
	o := orm.NewOrm()
	topics := make([]*Topic,0)
	qs := o.QueryTable("topic")
	var err error
	if isDesc {
		_,err = qs.OrderBy("-Created").All(&topics)
	}else {
		_,err = qs.All(&topics)
	}
	return topics,err
}

func GetTopic(tid string) (*Topic , error)  {
	tidNum , err := strconv.ParseInt(tid,10,64)
	if err != nil {
		return nil ,err
	}
	topic := new(Topic)
	o := orm.NewOrm()

	qs := o.QueryTable("topic")
	err = qs.Filter("id",tidNum).One(topic)
	if err != nil {
		return nil ,err
	}

	topic.Views ++
	_ ,err = o.Update(topic)
	return topic,err
}

func DelTopic(id string) error  {
	cid , err := strconv.ParseInt(id,10,64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Topic{Id:cid}
	_ ,err = o.Delete(cate)
	return err
}

func ModifyTopic(id ,title , content string) error  {
	tid ,err := strconv.ParseInt(id,10,64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{Id:tid}

	err = o.Read(topic)
	if err != nil{
		return err
	}
	topic.Title = title
	topic.Content = content
	topic.Created = time.Now()
	o.Update(topic)
	return nil
}



