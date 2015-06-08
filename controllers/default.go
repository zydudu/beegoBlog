package controllers

import (
	"beegoBlog/models"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	this.Data["blogs"] = models.GetAll()
	//log.Debug("debug", this.Data["blogs"])
	this.Layout = "layout.tpl"
	this.TplNames = "index.tpl"
}
