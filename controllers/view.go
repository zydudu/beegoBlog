package controllers

import (
	"beegoBlog/models"
	"github.com/astaxie/beego"
	"strconv"
)

type ViewController struct {
	beego.Controller
}

func (this *ViewController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	this.Data["Post"] = models.GetBlog(id)
	this.Layout = "layout.tpl"
	this.TplNames = "view.tpl"
}
