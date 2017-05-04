package routers

import (
	"bgbuildserver/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1/build", &controllers.MainController{}, "post:BuildPost;get:BuildGet")
	beego.Router("/api/v1/fproject", &controllers.MainController{}, "post:Fpost;get:Fget")
	beego.Router("/api/v1/fprojects", &controllers.MainController{}, "get:Fgets")
	beego.Router("/api/v1/sproject", &controllers.MainController{}, "post:Spost;get:Sget")
	beego.Router("/api/v1/sprojects", &controllers.MainController{}, "get:Sgets")
	beego.Router("/api/v1/replicas", &controllers.MainController{}, "post:Rpost;get:Rget")
	beego.Router("/api/v1/replicass", &controllers.MainController{}, "get:Rgets")
	beego.Router("/api/v1/service", &controllers.MainController{}, "post:Servicepost")

}
