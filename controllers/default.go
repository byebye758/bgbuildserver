package controllers

import (
	"bgbuildserver/models/build"
	"bgbuildserver/models/docker"
	f "bgbuildserver/models/fatherproject"
	r "bgbuildserver/models/replicas"
	"bgbuildserver/models/service"
	"bgbuildserver/models/software"
	s "bgbuildserver/models/sonproject"

	"encoding/json"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

var (
	Buildstatus = make(map[string]map[string]string)
)

type R struct {
	Status  string
	Url     string
	Buildid string
}

// func (c *MainController) Get() {
// 	c.Data["Website"] = "beego.me"
// 	c.Data["Email"] = "astaxie@gmail.com"
// 	c.TplName = "index.tpl"
// }

func (this *MainController) BuildPost() {
	bud := build.Appbuild{}
	bud.Init()
	beego.Info(bud)
	a := this.Ctx.Input.RequestBody
	//beego.Info(string(a))
	err := json.Unmarshal(a, &bud)
	if err != nil {
		beego.Error("Json to struct error")
	}

	//beego.Info(bud, bud.Projectname)
	id := bud.Projectname + bud.Buildid
	B := make(map[string]string)
	Buildstatus[id] = B
	Buildstatus[id]["Status"] = "1"
	go func() {

		err = bud.Buildpackage()
		if err != nil {
			beego.Error("build error")
			Buildstatus[id]["Status"] = "2"

		} else {
			url := "10.1.20.97:8080" + beego.AppConfig.String("staticurl") + "/" + bud.Projectname + "/" + bud.Buildid + "/" + bud.Staticpackagename

			Buildstatus[id]["Status"] = "0"
			Buildstatus[id]["Url"] = url

		}

	}()

	r := &R{
		Status:  Buildstatus[id]["Status"],
		Buildid: id,
	}
	b, err := json.Marshal(r)
	//beego.Info(string(b))
	if err != nil {
		beego.Error("post json Marshal  error")
	} else {

		this.Ctx.WriteString(string(b))

	}

}

func (this *MainController) BuildGet() {

	//a := State["k8s-test-goe9f70cc2-d963-42ee-8bd5-2cb12639e4fb"]
	bid := this.GetString("buildid")
	m := Buildstatus[bid]
	b, err := json.Marshal(m)

	if err != nil {
		beego.Error("post json Marshal  error")
	} else {
		//定义返回状态码
		this.Ctx.ResponseWriter.WriteHeader(201)
		this.Ctx.WriteString(string(b))
		beego.Info(Buildstatus)
		//beego.Info(a)
	}
}

func (this *MainController) DockerbuildPost() {
	dockerops := docker.Docker{}
	dockerops.Init()
	a := this.Ctx.Input.RequestBody
	err := json.Unmarshal(a, &dockerops)
	if err != nil {
		beego.Error("Json to struct error")
	}
}
func (this *MainController) DockerbuildGet() {

}

// func (this *MainController) Fget() {
// 	d := new(f.Fatherproject)
// 	d.Projectname = this.GetString("projectname")
// 	err := d.Init()
// 	if err != nil {
// 		beego.Error(err)

// 		this.Abort("401")
// 	}
// 	data, err := d.Get()
// 	if err != nil {
// 		beego.Error(err)

// 		this.Abort("401")
// 	}
// 	this.Ctx.WriteString(data)
// }
// func (this *MainController) Fgets() {
// 	d := new(f.Fatherproject)
// 	err := d.Init()
// 	if err != nil {
// 		beego.Error(err)
// 		this.Abort("401")
// 	}
// 	data, err := d.Gets()
// 	if err != nil {
// 		beego.Error(err)
// 		this.Abort("401")
// 	}
// 	this.Ctx.WriteString(data)
// }

// func (this *MainController) Fpost() {
// 	d := new(f.Fatherproject)

// 	body := this.Ctx.Input.RequestBody
// 	err := Postetcd(body, d)
// 	if err != nil {
// 		this.Abort("401")
// 	}
// 	beego.Info(d)
// 	this.Ctx.WriteString("ok")

// }
// func (this *MainController) Spost() {
// 	d := new(s.Sonproject)

// 	body := this.Ctx.Input.RequestBody
// 	err := Postetcd(body, d)
// 	if err != nil {
// 		this.Abort("401")
// 	}
// 	//beego.Info(d)
// 	this.Ctx.WriteString("ok")

// }

// func (this *MainController) Sget() {
// 	d := new(s.Sonproject)
// 	//d.Projectname = this.GetString("projectname")
// 	d.Sonprojectname = this.GetString("sprojectname")
// 	err := d.Init()
// 	if err != nil {
// 		beego.Error(err)

// 		this.Abort("401")
// 	}
// 	data, err := d.Get()
// 	if err != nil {
// 		beego.Error(err)

// 		this.Abort("401")
// 	}
// 	this.Ctx.WriteString(data)
// }
// func (this *MainController) Sgets() {
// 	d := new(s.Sonproject)
// 	err := d.Init()
// 	if err != nil {
// 		beego.Error(err)
// 		this.Abort("401")
// 	}
// 	data, err := d.Gets()
// 	if err != nil {
// 		beego.Error(err)
// 		this.Abort("401")
// 	}
// 	this.Ctx.WriteString(data)
// }

// func (this *MainController) Rpost() {
// 	d := new(r.Replicas)

// 	body := this.Ctx.Input.RequestBody
// 	err := Postetcd(body, d)
// 	if err != nil {
// 		this.Abort("401")
// 	}
// 	this.Ctx.WriteString("ok")
// }
// func (this *MainController) Rget() {
// 	d := new(r.Replicas)
// 	d.Sonprojectname = this.GetString("replicasname")
// 	err := d.Init()
// 	if err != nil {
// 		beego.Error(err)

// 		this.Abort("401")
// 	}
// 	data, err := d.Get()
// 	if err != nil {
// 		beego.Error(err)

// 		this.Abort("401")
// 	}
// 	this.Ctx.WriteString(data)
// }
// func (this *MainController) Rgets() {
// 	d := new(r.Replicas)
// 	err := d.Init()
// 	if err != nil {
// 		beego.Error(err)
// 		this.Abort("401")
// 	}
// 	data, err := d.Gets()
// 	if err != nil {
// 		beego.Error(err)
// 		this.Abort("401")
// 	}
// 	this.Ctx.WriteString(data)
// }
// func (this *MainController) Servicepost() {
// 	d := new(service.Service)

// 	body := this.Ctx.Input.RequestBody
// 	err := Postetcd(body, d)
// 	if err != nil {
// 		this.Abort("401")
// 	}
// 	this.Ctx.WriteString("ok")
// }
// func (this *MainController) Serviceget() {
// 	d := new(service.Service)
// 	d.Sonprojectname = this.GetString("servicename")
// 	err := d.Init()
// 	if err != nil {
// 		beego.Error(err)

// 		this.Abort("401")
// 	}
// 	data, err := d.Get()
// 	if err != nil {
// 		beego.Error(err)

// 		this.Abort("401")
// 	}
// 	this.Ctx.WriteString(data)
// }
// func (this *MainController) Servicegets() {
// 	d := new(service.Service)
// 	err := d.Init()
// 	if err != nil {
// 		beego.Error(err)
// 		this.Abort("401")
// 	}
// 	data, err := d.Gets()
// 	if err != nil {
// 		beego.Error(err)
// 		this.Abort("401")
// 	}
// 	this.Ctx.WriteString(data)
// }
// func (this *MainController) Softwarepost() {
// 	d := new(software.Software)

// 	body := this.Ctx.Input.RequestBody
// 	err := Postetcd(body, d)
// 	if err != nil {
// 		this.Abort("401")
// 	}
// 	this.Ctx.WriteString("ok")
// }

// func (this *MainController) Softwareget() {
// 	d := new(software.Software)
// 	d.Sonprojectname = this.GetString("softwarename")
// 	err := d.Init()
// 	if err != nil {
// 		beego.Error(err)

// 		this.Abort("401")
// 	}
// 	data, err := d.Get()
// 	if err != nil {
// 		beego.Error(err)

// 		this.Abort("401")
// 	}
// 	this.Ctx.WriteString(data)
// }
// func (this *MainController) Softwaregets() {
// 	d := new(software.Software)
// 	err := d.Init()
// 	if err != nil {
// 		beego.Error(err)
// 		this.Abort("401")
// 	}
// 	data, err := d.Gets()
// 	if err != nil {
// 		beego.Error(err)
// 		this.Abort("401")
// 	}
// 	this.Ctx.WriteString(data)
// }
func Postetcd(body []byte, st interface{}) error {
	//d := new(st)
	err := json.Unmarshal(body, &st)
	if err != nil {
		beego.Error(err)
		return err
	}
	switch d := st.(type) {
	case *f.Fatherproject:
		err = d.Init()
		if err != nil {
			beego.Error(err)
			return err
		}
		err = d.Put(d)
		if err != nil {
			beego.Error(err)
			return err
		}
	case *s.Sonproject:
		err = d.Init()
		if err != nil {
			beego.Error(err)
			return err
		}
		err = d.Put(d)
		if err != nil {
			beego.Error(err)
			return err
		}
	case *r.Replicas:
		err = d.Init()
		if err != nil {
			beego.Error(err)
			return err
		}
		err = d.Put(d)
		if err != nil {
			beego.Error(err)
			return err
		}
	case *service.Service:
		err = d.Init()
		if err != nil {
			beego.Error(err)
			return err
		}
		err = d.Put(d)
		if err != nil {
			beego.Error(err)
			return err
		}
	case *software.Software:
		err = d.Init()
		if err != nil {
			beego.Error(err)
			return err
		}
		err = d.Put(d)
		if err != nil {
			beego.Error(err)
			return err
		}

	}

	return nil

}
