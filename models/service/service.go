package service

import (
	"bgbuildserver/models/etcd"
)

const (
	keypath = "/autocloud/service/"
)

type Service struct {
	Projectname    string
	Sonprojectname string
	Servicename    string
	Serviceswitch  bool    `json:"Serviceswitch,string,omitempty"`
	Servicetype    int64   `json:"Servicetype,string,omitempty"`
	ServicePorts   []int64 `json:"Servicetype,string,omitempty"`
	etcd.Etcd
}

func (s *Service) Init() error {
	s.Servicename = s.Sonprojectname
	s.Keypath = keypath
	s.Key = s.Keypath + s.Servicename
	return nil
}
