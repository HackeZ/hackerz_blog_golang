package routers

import (
	"github.com/astaxie/beego"
	"hackerz_blog_golang/controllers"
)

func init() {
	// 注册路由
	// 1. 这是最常见的路由，固定路由
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/reply", &controllers.ReplyController{})

	// 2. 这是指定路由
	beego.Router("/reply/add", &controllers.ReplyController{}, "POST:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "GET:Delete")

	// 3. 智能路由，但凡提交的找不到对应上面的路由，都会提交到智能路由...
	beego.AutoRouter(&controllers.TopicController{})
}
