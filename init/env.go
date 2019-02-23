//add by Chengshipeng
//if we have the same init enviorment
//配置环境初始化
package init

import (
	"github.com/astaxie/beego"
)

func InitEnv() {
	runmode := beego.AppConfig.String("runmode")
	if runmode == "dev" {
		//beego.SetStaticPath("/static", "static")
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
}
