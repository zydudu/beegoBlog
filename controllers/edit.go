package controllers

import (
	"beegoBlog/models"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type EditController struct {
	beego.Controller
}

func (this *EditController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	this.Data["Post"] = models.GetBlog(id)
	log.Debug("debug", this.Data["Post"])
	this.Layout = "layout.tpl"
	this.TplNames = "edit.tpl"
}

func (this *EditController) Post() {
	inputs := this.Input()
	var blog models.Blog
	blog.Id, _ = strconv.Atoi(inputs.Get("id"))
	blog.Title = inputs.Get("title")
	blog.Content = inputs.Get("content")
	blog.Modified = time.Now()
	models.SaveBlog(blog)
	this.Ctx.Redirect(302, "/")
}
