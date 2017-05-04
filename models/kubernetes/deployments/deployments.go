package deployments

import (
	"flag"
	"fmt"
	//"time"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/astaxie/beego"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/unversioned"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/apis/extensions/v1beta1"
	"k8s.io/client-go/tools/clientcmd"
)

func init() {
	kubeconfigpath = beego.AppPath + "/conf/kubeconfig"
	ubeconfig := flag.String("kubeconfig", kubeconfigpath, "absolute path to the kubeconfig file")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	Clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}
