package sonproject

import (
	"bgbuildserver/models/etcd"
	// "github.com/astaxie/beego"
)

const (
	keypath = "/autocloud/sonproject/"
)

type Sonproject struct {
	Projectname    string
	Sonprojectname string
	etcd.Etcd
}

func (s *Sonproject) Init() error {
	s.Keypath = keypath
	s.Key = s.Keypath + s.Sonprojectname
	return nil
}
