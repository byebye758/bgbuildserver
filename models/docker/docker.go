package docker

import (
	"archive/tar"
	"bytes"
	"github.com/astaxie/beego"
	docker "github.com/fsouza/go-dockerclient"
	//"strings"

	"github.com/golibs/uuid"
	"time"
)

type Docker struct {
	Projectname         string
	Projectid           string
	ImageFrom           string //10.10.30.169/library/tomcat:v02
	PackageUrl          string //10.1.20.97:8080/staticdownload/k8s-test-go/18a97097-45b7-4d7d-9e5e-e037c6dab2a2/snpushandroid.tar.gz
	ImageName           string //tomcat
	ImageTag            string //vtest
	RegistryAddress     string //10.10.30.169
	RegistryTag         string //library
	RegistryUsername    string //admin
	RegistryUserpasswd  string //admin
	Dockerfile          []byte
	DockerServeraddress string //tcp://10.10.30.244:2375
	Buildid             string
}

//{"ImageFrom":"10.10.30.169/library/tomcat:v02","PackageUrl":"//10.1.20.97:8080/staticdownload/k8s-test-go/18a97097-45b7-4d7d-9e5e-e037c6dab2a2/snpushandroid.tar.gz","Cmd":[]string{"/usr/local/tomcat/bin/catalina.sh","run"},"ImageName":"tomcat","ImageTag":"vtest","RegistryAddress":"10.10.30.169","RegistryTag":"admin","RegistryUserpasswd":"admin","DockerServeraddress":"tcp://10.10.30.244:2375"}
func (d *Docker) BuildPush() error {
	d.Dockerfile = []byte("FROM " + d.ImageFrom + "\n" + "ADD " + d.PackageUrl + "\n")

	client, err := docker.NewClient(d.DockerServeraddress)
	if err != nil {
		beego.Error(err, "Docerserver connect error")
		return err
	}
	t := time.Now()
	inputbuf, outputbuf := bytes.NewBuffer(nil), bytes.NewBuffer(nil)
	tr := tar.NewWriter(inputbuf)
	tr.WriteHeader(&tar.Header{Name: "Dockerfile", Size: int64(len(d.Dockerfile)), ModTime: t, AccessTime: t, ChangeTime: t})
	tr.Write(d.Dockerfile)
	tr.Close()
	opts := docker.BuildImageOptions{
		Name:                d.RegistryAddress + "/" + d.RegistryTag + "/" + d.ImageName + ":" + d.ImageTag,
		InputStream:         inputbuf,
		OutputStream:        outputbuf,
		ForceRmTmpContainer: true,
		RmTmpContainer:      true,
		NoCache:             true,
	}
	if err := client.BuildImage(opts); err != nil {
		beego.Error(err)
		return err

	}

	auth := docker.AuthConfiguration{
		Username:      d.RegistryUsername,
		Password:      d.RegistryUserpasswd,
		ServerAddress: d.RegistryAddress,
	}
	pushopts := docker.PushImageOptions{
		Name:     d.RegistryAddress + "/" + d.RegistryTag + "/" + d.ImageName,
		Tag:      d.ImageTag,
		Registry: d.RegistryAddress + "/" + d.RegistryTag + "/",
	}
	if err := client.PushImage(pushopts, auth); err != nil {
		beego.Error(err)
		return err
	}
	return nil

}
func (d *Docker) Init() {
	d.Buildid = uuid.Rand().Hex()
}
