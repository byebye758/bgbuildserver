package controllers

import (
	f "bgbuildserver/models/fatherproject"
	"github.com/astaxie/beego"
)

func (this *MainController) Fget() {
	d := new(f.Fatherproject)
	d.Projectname = this.GetString("projectname")
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
func (this *MainController) Fgets() {
	d := new(f.Fatherproject)
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

func (this *MainController) Fpost() {
	d := new(f.Fatherproject)

	body := this.Ctx.Input.RequestBody
	err := Postetcd(body, d)
	if err != nil {
		this.Abort("401")
	}
	beego.Info(d)
	this.Ctx.WriteString("ok")

}
