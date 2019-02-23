package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"booksmanagementSys/models"
)

func DbInit() {
	/*mysql
	//DSN数据源字符串：用户名:密码@协议(地址:端口)/数据库?参数=参数值
		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.RegisterDataBase("default", "mysql", "root:abcd1234@tcp(10.2.45.55:3306)/libmanage?charset=utf8")
		orm.RunSyncdb("default", false, true) //同步建表
	dsn数据源字符串：用户名：密码@协议（主机:端口/数据库？参数=参数值*/
	mdb_host := beego.AppConfig.String("mdb.host")
	mdb_port := beego.AppConfig.String("mdb.port")
	mdb_user := beego.AppConfig.String("mdb.user")
	mdb_password := beego.AppConfig.String("mdb.password")
	mdb_name := beego.AppConfig.String("mdb.name")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", mdb_user, mdb_password, mdb_host,mdb_port, mdb_name)
	fmt.Println("dsn:", dsn)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)
	/*初始化model*/
	orm.RegisterModel(new(models.Book))
	orm.RunSyncdb("default", false, true)
}
