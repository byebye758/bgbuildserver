package controllers

import (
	"bgbuildserver/models/service"
	"github.com/astaxie/beego"
)

func (this *MainController) Servicepost() {
	d := new(service.Service)

	body := this.Ctx.Input.RequestBody
	err := Postetcd(body, d)
	if err != nil {
		this.Abort("401")
	}
	this.Ctx.WriteString("ok")
}
func (this *MainController) Serviceget() {
	d := new(service.Service)
	d.Sonprojectname = this.GetString("servicename")
	err := d.Init()
	if err != nil {
		beego.Error(err)

		this.Abort("401")
	}
	data, err := d.Get()
	if err != nil {
		beego.Error(err)

		this.Abort("401")
	}
	this.Ctx.WriteString(data)
}
func (this *MainController) Servicegets() {
	d := new(service.Service)
	err := d.Init()
	if err != nil {
		beego.Error(err)
		this.Abort("401")
	}
	data, err := d.Gets()
	if err != nil {
		beego.Error(err)
		this.Abort("401")
	}
	this.Ctx.WriteString(data)
}
