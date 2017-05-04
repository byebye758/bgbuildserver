package etcd

import (
	//"fmt"
	//"bgbuildserver/models/fatherproject"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
	"time"
)

const (
	Timout = 5 * time.Second
)

var (
	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: Timout,
	}
)

type Etcd struct {
	Key     string `json:"-"`
	Keypath string `json:"-"`
}

func EtcdPut(key, value string) error {
	cli, err := clientv3.New(config)
	if err != nil {
		beego.Error(err)
	}

	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), Timout)
	_, err = cli.Put(ctx, key, value)
	cancel()
	if err != nil {
		beego.Error(err)
		return err
	} else {
		return nil
	}

}

func Etcdgetkeys(key string) (kvmap map[string][]byte, jss string, err error) {
	cli, err := clientv3.New(config)
	if err != nil {
		beego.Error(err)
	}

	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), Timout)
	rest, err := cli.Get(ctx, key, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	cancel()
	if err != nil {
		beego.Error(err)
		return nil, "", err
	}
	kvmap = make(map[string][]byte)
	jsmap := make(map[string]string)
	for _, v := range rest.Kvs {
		kvmap[string(v.Key)] = v.Value
		jsmap[string(v.Key)] = string(v.Value)
	}
	data, err := json.Marshal(&jsmap)
	if err != nil {
		return nil, "", err
		beego.Error(err)
	}
	err = nil
	jss = string(data)
	return kvmap, jss, err
}
func Etcdgetkey(key string) (jsb []byte, jss string, err error) {
	cli, err := clientv3.New(config)
	if err != nil {
		beego.Error(err)
	}

	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), Timout)
	rest, err := cli.Get(ctx, key)
	cancel()

	if err != nil {
		beego.Error(err)
	}
	jsb = rest.Kvs[0].Value
	jss = string(jsb)
	beego.Info(string(jsb))
	err = nil
	return jsb, jss, err
}

//json 类型返回
// func Etcdgetkeysjson(keypath string) (string, error) {
// 	b, err := Etcdgetkeys(keypath)
// 	if err != nil {
// 		return "", err
// 		beego.Error(err)
// 	}
// 	j := make(map[string]string)
// 	for k, v := range b {
// 		j[k] = string(v)
// 	}
// 	data, err := json.Marshal(j)
// 	if err != nil {
// 		return "", err
// 		beego.Error(err)
// 	}

// 	return string(data), nil
// }

// func Etcdgetkeyjson(key string) (string, error) {
// 	b, err := Etcdgetkey(key)
// 	if err != nil {
// 		return "", err
// 		beego.Error(err)
// 	}
// 	data, err := json.Marshal(b)
// 	if err != nil {
// 		return "", err
// 		beego.Error(err)
// 	}

// 	return string(data), nil
// }

func (e *Etcd) Gets() (string, error) {
	_, data, err := Etcdgetkeys(e.Keypath)
	if err != nil {
		return "", err

		beego.Error(err)
	}
	return data, nil

}

func (e *Etcd) Put(s interface{}) error {

	b, err := json.Marshal(s)
	if err != nil {

		beego.Error(err)
		return err
	}

	err = EtcdPut(e.Key, string(b))
	if err != nil {

		beego.Error(err)
		return err
	}
	return nil
}

func (e *Etcd) Get() (string, error) {
	_, data, err := Etcdgetkey(e.Key)
	if err != nil {
		return "", err

		beego.Error(err)
	}
	return data, nil
}
