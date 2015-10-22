package controllers

import (
	"hackerz_blog_golang/models"
	"strings"

	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.Data["IsLogin"] = checkLogin(c.Ctx)
	c.Data["IsTopic"] = true
	c.TplNames = "topic.html"

	// 获取所有文章
	// 参数：分类、标签、排序
	topics, err := models.GetAllTopics("", "", false)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["Topics"] = topics
	}
}

func (c *TopicController) Post() {
	if !checkLogin(c.Ctx) {
		c.Redirect("/", 302)
		return
	}

	tid := c.Input().Get("tid")
	title := c.Input().Get("title")
	category := c.Input().Get("category")
	label := c.Input().Get("label")
	content := c.Input().Get("content")

	var err error
	// 这里判断到底是添加文章还是修改文章
	// PS：这样写会被恶意POST数据，需要修改
	if len(tid) == 0 {
		err = models.AddTopic(title, category, label, content)
	} else {
		err = models.ModifyTopic(tid, title, category, label, content)
	}

	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/topic", 302)
}

// 添加文章
func (c *TopicController) Add() {
	c.TplNames = "topic_add.html"
}

// 查看文章详情
func (c *TopicController) View() {
	c.TplNames = "topic_view.html"

	// 获取文章数据
	topic, err := models.GetTopic(c.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	// 先将文章内容数据丢出给前台，就不论是否获取评论成功与否，文章都能显示
	c.Data["Topic"] = topic
	c.Data["Label"] = strings.Split(topic.Label, " ")
	c.Data["Tid"] = c.Ctx.Input.Param("0")

	// 获取评论数据
	replies, err := models.GetAllReplies(c.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		return
	}

	c.Data["Replies"] = replies
	c.Data["IsLogin"] = checkLogin(c.Ctx)
}

// 修改文章
func (c *TopicController) Modify() {
	c.TplNames = "topic_modify.html"

	if !checkLogin(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}

	tid := c.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Data["Topic"] = topic
	c.Data["Tid"] = tid
}

// 删除文章
func (c *TopicController) Delete() {
	if !checkLogin(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}

	err := models.DeleteTopic(c.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/topic", 302)
}
