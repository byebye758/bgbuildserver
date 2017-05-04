package controllers

import (
	r "bgbuildserver/models/replicas"
	"github.com/astaxie/beego"
)

func (this *MainController) Rpost() {
	d := new(r.Replicas)

	body := this.Ctx.Input.RequestBody
	err := Postetcd(body, d)
	if err != nil {
		this.Abort("401")
	}
	this.Ctx.WriteString("ok")
}
func (this *MainController) Rget() {
	d := new(r.Replicas)
	d.Sonprojectname = this.GetString("replicasname")
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
func (this *MainController) Rgets() {
	d := new(r.Replicas)
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
