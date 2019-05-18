package main

import (
	_ "ownergit/booksmanagement/routers"

	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"
	"ownergit/booksmanagement/utils"
	"github.com/astaxie/beego/orm"
	"port_scan"
	"flag"
	"strconv"
)

func init() {
	utils.DbInit()
}

func main() {
	// 增加随机端口、命令行不同模式启动
	var (
		svrPort string
	)
	port_scan.FlagRunmode()
	flag.Parse()
	portInt := port_scan.RandomPort()
	svrPort = strconv.Itoa(portInt)
	if len(svrPort) <= 0 {
		svrPort = "10010"
	}
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Info("服务已启动.......监听端口：",svrPort)
	beego.Run(":" + svrPort)
}
