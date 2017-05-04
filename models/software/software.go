package software

import (
	"bgbuildserver/models/etcd"
)

const (
	keypath = "/autocloud/software/"
)

type Software struct {
	Sonprojectname     string
	Projectname        string
	Softwarename       string
	Softwareimages     string
	Softwareurl        string
	Softwarepath       string
	Softwareconfigpath string
	Softwaresecretpath string
	Softwarecmd        []string
	//Storage
	//Storagepath
	etcd.Etcd
}

func (s *Software) Init() error {
	s.Keypath = keypath
	s.Key = s.Keypath + s.Softwarename
	return nil
}
