package main

import (
	_ "bgbuildserver/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.SetStaticPath(beego.AppConfig.String("staticurl"), beego.AppConfig.String("staticpath"))
	//beego.Error(beego.AppConfig.String("staticpath"))
	beego.BConfig.Log.AccessLogs = true
	beego.BConfig.Log.FileLineNum = true
	//beego.BConfig.Log.Outputs = map[string]{}
	beego.SetLogger("file", `{"filename":"test.log"}`)
	beego.Run()
}
