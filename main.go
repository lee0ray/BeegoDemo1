package main

import (
	"github.com/astaxie/beego"
	"myfirst/commands"
	_ "myfirst/models"
	_ "myfirst/routers"
)

func main() {
	commands.Bootstrap()
	beego.Run()
}

