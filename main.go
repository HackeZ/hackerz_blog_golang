package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"hackerz_blog_golang/models"
	_ "hackerz_blog_golang/routers"
)

func init() {
	// 初始化数据库
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
