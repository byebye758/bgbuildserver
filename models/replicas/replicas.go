package replicas

import (
	"bgbuildserver/models/etcd"
)

const (
	keypath = "/autocloud/replicas/"
)

type Replicas struct {
	Projectname    string
	Sonprojectname string
	Replicasname   string
	Min            int64 `json:"Min,string,omitempty"`
	Max            int64 `json:"Max,string,omitempty"`
	Autopolicy     bool  `json:"Autopolicy,string,omitempty"`
	Autotype       int64 `json:"string,omitempty"`
	etcd.Etcd
}

func (r *Replicas) Init() error {
	r.Replicasname = r.Sonprojectname
	r.Keypath = keypath
	r.Key = r.Keypath + r.Replicasname
	return nil
}
