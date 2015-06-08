package main

import (
	_ "beegoBlog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init() {
	log := logs.NewLogger(10000)
	log.SetLogger("console", "")
}
func main() {
	beego.Run()
}
