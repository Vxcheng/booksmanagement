//add by xjy 2018.07.05
package models

import (
	"errors"
	"fmt"
	"time"

	//log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//MYSQL
type DBConfig struct {
	Host         string
	Port         string
	Database     string
	Username     string
	Password     string
	MaxIdleConns int //最大空闲连接
	MaxOpenConns int //最大连接数
}

//USER
type User struct {
	Id         int64 `orm:"column(uid);pk;auto;default(1)"`
	Psw        string
	Uname      string    `orm:"size(100);unique"`
	CreateTime time.Time `orm:"column(create_time);null"` //`orm:"index"`
	LastTime   time.Time `orm:"column(last_time);null"`   //`orm:"index"`
}

type UserManager struct {
	DBConf *DBConfig
}

func NewUserManager(dbConf *DBConfig) *UserManager {
	mgr := &UserManager{
		DBConf: dbConf,
	}
	mgr.initDB() //初始化orm
	return mgr
}

// 初始化db，注册默认数据库，同时将实体模型也注册上去
func (mgr *UserManager) initDB() {
	//registe driver
	orm.RegisterDriver("mysql", orm.DRMySQL)
	ds := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", mgr.DBConf.Username, mgr.DBConf.Password, mgr.DBConf.Host, mgr.DBConf.Port, mgr.DBConf.Database)
	//log.Infof("datasource=[%s]", ds)

	err := orm.RegisterDataBase("default", "mysql", ds, mgr.DBConf.MaxIdleConns, mgr.DBConf.MaxOpenConns)
	if err != nil {
		//log.Error(err)
	}
	// 注册数据库
	orm.RegisterModel(new(User))
	// 自动建表
	orm.RunSyncdb("default", false, true)
}

/****************************************************************
【函数】：GetTableName()
【功能】:获取user的表名
【传入参数】:
【输出参数】:string
【创建时间】：2018.07.06 by xjy
*******************************************************************/
func (user *User) GetTableName() string {

	return "database.user"
}

/****************************************************************
【函数】：AddUser()
【功能】:增加user，配合注册用户使用
【传入参数】:*User
【输出参数】:(int64, error)uid及错误信息
【创建时间】：2018.07.06 by xjy
*******************************************************************/
func AddUser(user *User) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(user)
	//log.Infof("Insert user:%+v", user)
	if err != nil {
		//log.Error("AddUser error:", err.Error())
		return 0, err
	}
	return id, nil
}

/****************************************************************
【函数】：UpdateUserById()
【功能】:通过uid更新用户密码
【传入参数】:*User
【输出参数】:error错误信息
【创建时间】：2018.07.06 by xjy
*******************************************************************/
func UpdateUserById(user *User) error {
	o := orm.NewOrm()
	v := User{Id: user.Id}
	// ascertain id exists in the database
	if err := o.Read(&v); err == nil {
		if num, err := o.Update(user, "Psw", "LastTime"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}

	return nil
}

/****************************************************************
【函数】：GetUserByFilter()
【功能】:通过uid或者uname查询user记录
【传入参数】:*User
【输出参数】:error错误信息
【创建时间】：2018.07.06 by xjy
*******************************************************************/
func GetUserByFilter(field string, user interface{}) ([]orm.Params, error) {

	var (
		num  int64
		maps []orm.Params
		err  error
	)
	o := orm.NewOrm()
	switch field {
	case "uname":
		num, err = o.Raw("select * FROM user WHERE uname=? ", user).Values(&maps)
	case "uid":

		num, err = o.Raw("select * FROM user WHERE uid=? ", user).Values(&maps)
	default:
		return nil, errors.New("field type error")

	}
	if err == nil && num > 0 {
		return maps, nil
	}
	return nil, err
}

/****************************************************************
【函数】：VerifyUser()
【功能】:根据传入的user（password&uid password&uname）确认是否存在对应user记录
【传入参数】:user *User
【输出参数】:([]orm.Params, error)
【创建时间】：2018.07.06 by xjy
*******************************************************************/
func VerifyUser(user *User) ([]orm.Params, error) {
	//orm.Debug = true
	var (
		num  int64
		maps []orm.Params
		err  error
	)
	o := orm.NewOrm()

	if user.Id != 0 {
		num, err = o.Raw("select * FROM user WHERE uid=? and psw=? ", user.Id, user.Psw).Values(&maps)

	} else if user.Uname != "" {
		num, err = o.Raw("select * FROM user WHERE uname=? and psw=? ", user.Uname, user.Psw).Values(&maps)
	} else {
		return nil, errors.New("uid and uname not found")
	}
	//beego.Debug("row nums: ", num)
	if err == nil && num > 0 {
		return maps, nil
	}

	return nil, err
}

/****************************************************************
【函数】：DeleteUser
【功能】:删除uid或者username对应的记录
【传入参数】:uid int64
【输出参数】:(interface{}, error)
【创建时间】：2018.07.06 by xjy
*******************************************************************/
func DeleteUser(uid int64) (interface{}, error) {
	o := orm.NewOrm()
	user := User{Id: uid}
	_, err := o.Delete(&user)
	if err != nil {
		//log.Error("QueryUser error:", err.Error())
		return nil, err
	}
	return user, nil
}
