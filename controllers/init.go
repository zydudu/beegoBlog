package controllers

import (
	"github.com/astaxie/beego/logs"
)

var log = logs.NewLogger(10000)

func init() {
	log.SetLogger("console", "")
}
