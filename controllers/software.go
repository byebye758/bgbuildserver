package controllers

import (
	"bgbuildserver/models/software"
	"github.com/astaxie/beego"
)

func (this *MainController) Softwarepost() {
	d := new(software.Software)

	body := this.Ctx.Input.RequestBody
	err := Postetcd(body, d)
	if err != nil {
		this.Abort("401")
	}
	this.Ctx.WriteString("ok")
}

func (this *MainController) Softwareget() {
	d := new(software.Software)
	d.Sonprojectname = this.GetString("softwarename")
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
func (this *MainController) Softwaregets() {
	d := new(software.Software)
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
