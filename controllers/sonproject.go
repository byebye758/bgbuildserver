package controllers

import (
	s "bgbuildserver/models/sonproject"
	"github.com/astaxie/beego"
)

func (this *MainController) Spost() {
	d := new(s.Sonproject)

	body := this.Ctx.Input.RequestBody
	err := Postetcd(body, d)
	if err != nil {
		this.Abort("401")
	}
	//beego.Info(d)
	this.Ctx.WriteString("ok")

}

func (this *MainController) Sget() {
	d := new(s.Sonproject)
	//d.Projectname = this.GetString("projectname")
	d.Sonprojectname = this.GetString("sprojectname")
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
func (this *MainController) Sgets() {
	d := new(s.Sonproject)
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
	beego.Info(beego.AppPath)

	this.Ctx.WriteString(data)
}
