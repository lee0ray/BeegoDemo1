package commands

import (
	"github.com/astaxie/beego"
)

func Bootstrap() {
	beego.AddFuncMap("PrePage", PrePage)
	beego.AddFuncMap("NextPage", NextPage)
}
func PrePage(data int) int {
	data--
	if data<=0{
		data =1
	}
	return data
}
func NextPage(data int) int {
	data++
	return data
}
