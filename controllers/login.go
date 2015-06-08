package controllers

import (
	//"beegoBlog/models"
	"github.com/astaxie/beego"
	//"time"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.Layout = "layout.tpl"
	this.TplNames = "login.tpl"
}

func (this *LoginController) Post() {
	//inputs := this.Input()

	//username := inputs.Get("username")
	//password := inputs.Get("password")

	this.Ctx.Redirect(302, "/new")
}
