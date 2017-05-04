package build

import (
	//"encoding/json"
	"encoding/xml"
	//"fmt"

	"github.com/astaxie/beego"
	"github.com/golibs/uuid"
	"github.com/mholt/archiver"
	"io/ioutil"
	"os"
	"os/exec"
)

var (
	staticpath = beego.AppConfig.String("staticpath")
	bbuildpath = beego.AppConfig.String("bbuildpath")
)

type Appbuild struct {
	Projectname       string
	Projectid         string
	Codetools         string //"svn" "git"
	Codeurl           string //http://tools.scinan.com.cn:88/svn/platform/branches/SNPushServer_develop
	Codeuser          string
	Codepasswd        string
	Buildtools        string //"mvn"
	Buildpar          string
	Buildpath         string
	BBuildpath        string
	Packagename       string
	Packagetype       string
	Packagepath       string
	Staticpackagename string
	Staticpackagepath string
	Staticpath        string
	Buildid           string
}

func (a *Appbuild) Buildpackage() error {
	//app := Appbuild{Staticpath: staticpath}
	//a.statusid = 1

	switch a.Codetools {
	case "svn":
		err := a.SVN()
		if err != nil {
			beego.Error("svn error")
			return err
		}
		switch a.Buildtools {
		case "mvn":
			err := a.MVN()
			if err != nil {
				beego.Error("mvn error")
				return err
			}
		}
	case "git":
		//app.GIT()
	}
	//a.statusid = 0
	return nil

}

func (a *Appbuild) Init() {
	//buildid
	a.Buildid = uuid.Rand().Hex()
	beego.Info(a.Buildid)
	a.Buildpath = bbuildpath + "/" + a.Projectname + "/" + a.Buildid
	a.Staticpath = staticpath

}
func (a *Appbuild) SVN() error {
	cmd := exec.Command("sh", "-c", "cd "+a.Buildpath+" && svn  update --username="+a.Codeuser+" --password="+a.Codepasswd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		cmd := exec.Command("sh", "-c", "svn co  "+a.Codeurl+" "+a.Buildpath)
		output, err := cmd.CombinedOutput()
		beego.Error("co", err, string(output))
		return err

	}
	beego.Error(string(output), err)
	return err
}

/*使用 MVN 进行编译，根据pom.xml 获取 build 生成的 package 名
package path 压缩数据包 并将包 传送到静态文件目录
*/
func (a *Appbuild) MVN() error {
	//mvn
	cmd := exec.Command("sh", "-c", "cd "+a.Buildpath+" &&  mvn "+a.Buildpar)
	output, err := cmd.CombinedOutput()
	beego.Info(string(output))
	if err != nil {
		beego.Error("Mvn  build Error")
		return err
	}

	//解析pom
	type Project struct {
		Name      string `xml:"name"`
		Packaging string `xml:"packaging"`
	}

	file, err := os.Open(a.Buildpath + "/pom.xml")

	if err != nil {
		beego.Error("open xml file error")
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		beego.Error("read file stream error")
		return err
	}

	project := Project{}
	err = xml.Unmarshal(data, &project)
	if err != nil {
		beego.Error("read file stream error")
		return err
	}
	a.Packagetype = project.Packaging
	switch a.Packagetype {
	case "jar":
		a.Packagename = project.Name
		a.Packagepath = a.Buildpath + "/target/" + a.Packagename
	case "war":
		a.Packagename = project.Name + a.Packagetype
		a.Packagepath = a.Buildpath + "/target/" + a.Packagename
	}

	a.Staticpackagename = a.Packagename + ".tar.gz"
	a.Staticpackagepath = a.Staticpath + "/" + a.Projectname + "/" + a.Buildid + "/" + a.Staticpackagename

	err = os.MkdirAll(a.Staticpath+"/"+a.Projectname+"/"+a.Buildid, 0755)
	if err != nil {
		beego.Error("Mkdir Error", err)
		return err
	} else {
		beego.Info("Mkdir " + a.Staticpath + "/" + a.Projectname + "/" + a.Buildid + "Ok")
	}

	err = archiver.TarGz.Make(a.Staticpackagepath, []string{a.Packagepath})
	if err != nil {
		beego.Error("TarGz Error")
		return err
	} else {
		beego.Info("TgGz " + a.Staticpackagepath + "Ok")

	}
	return nil
}
