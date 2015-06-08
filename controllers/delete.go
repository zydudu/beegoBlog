package controllers

import (
	"beegoBlog/models"
	"github.com/astaxie/beego"
	"strconv"
)

type DeleteController struct {
	beego.Controller
}

func (this *DeleteController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	log.Debug("debug", id)
	blog := models.GetBlog(id)
	this.Data["Post"] = blog
	models.DelBlog(blog)
	this.Ctx.Redirect(302, "/")
}
