//add by Chengshipeng
//if we have the same init enviorment
//数据库初始化
package init

import (
	_ "fmt"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/gogather/com"
)

func InitSql() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/mybeego")
	//orm.RegisterDataBase("dbB", "mysql", "root:123456@tcp(127.0.0.1:3306)/syslog")

}
