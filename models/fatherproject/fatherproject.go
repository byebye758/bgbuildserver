package fatherproject

import (
	"bgbuildserver/models/etcd"
	// //"bgbuildserver/models/sonproject"
	// "encoding/json"
	// "github.com/astaxie/beego"
)

const (
	keypath = "/autocloud/fatherproject/"
)

type Fatherproject struct {
	Projectname       string
	Fatherprojectname string
	etcd.Etcd
}

func (f *Fatherproject) Init() error {

	f.Fatherprojectname = "F" + f.Projectname
	f.Keypath = keypath
	f.Key = f.Keypath + f.Fatherprojectname

	return nil
}
