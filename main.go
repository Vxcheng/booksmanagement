package main

import (
	_ "booksmanagementSys/routers"

	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"
	"booksmanagementSys/utils"
)

func init() {
	utils.DbInit()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
