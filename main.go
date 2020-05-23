package main

import (
	"github.com/astaxie/beego"
	_ "myfirst/models"
	_ "myfirst/routers"
)

func main() {
	beego.Run()
}
