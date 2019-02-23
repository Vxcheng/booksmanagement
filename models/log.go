//add by Chengshipeng
//rizhixitong
package models

import (
	"log"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type JsonTime time.Time

type Log struct {
	LogId          int64     `json:"logId" orm:"pk;auto;column(logId)"`           //主键，自动增长
	UserId         int64     `json:"userId" orm:"column(userId)"`                 //用户唯一id
	BookId         int64     `json:"bookId" orm:"column(bookId)"`                 //书籍唯一id
	OperateContent string    `json:"operateContent" orm:"column(operateContent)"` //操作内容
	OperateDate    time.Time `json:"operateDate" orm:"column(operateDate)"`       //操作日期
}

type Log2 struct {
	LogId          int64    `json:"logId" orm:"pk;auto;column(logId)"`          //主键，自动增长
	UserId         int64    `json:"userId"`                                     //用户唯一id
	BookId         int64    `json:"bookId"orm:"column(bookId)"`                 //书籍唯一id
	OperateContent string   `json:"operateContent"orm:"column(operateContent)"` //操作内容
	OperateDate    JsonTime `json:"operateDate" orm:"column(operateDate)"`      //操作日期
}

type Log3 struct {
	LogId          int64     `json:"logId" orm:"pk;auto;column(logId)"`
	Username       string    `json:"username"orm:"column(username)"`
	Bookname       string    `json:"bookName"orm:"column(bookName)"`
	OperateContent string    `json:"operateContent"orm:"column(operateContent)"`
	OperateDate    time.Time `json:"operateDate" orm:"column(operateDate)"` //操作日期
}

type Log4 struct {
	LogId          int64    `json:"logId" orm:"pk;auto;column(logId)"`
	Username       string   `json:"username"orm:"column(username)"`
	Bookname       string   `json:"bookName"orm:"column(bookName)"`
	OperateContent string   `json:"operateContent"orm:"column(operateContent)"`
	OperateDate    JsonTime `json:"operateDate" orm:"column(operateDate)"` //操作日期
}

func init() {
	orm.RegisterModel(new(Log))
}

func (this JsonTime) MarshalJSON() ([]byte, error) {
	var tamp = strconv.FormatInt((time.Time(this).UnixNano() / 1000000), 10)
	return []byte(tamp), nil
}

func (this *Log) TableName() string {
	return "log"
}

func GetAllLogs() []*Log3 {
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("log.logId", "libuser.username", "book.bookName", "log.operateContent", "log.operateDate").From("log").InnerJoin("libuser").On("log.userId=libuser.id").InnerJoin("book").On("log.bookId=book.bookId")
	// 导出 SQL 语句
	sql := qb.String()
	o2 := orm.NewOrm()
	o2.Using("default")
	var log []*Log3
	//q := o2.QueryTable("log")
	o2.Raw(sql).QueryRows(&log)
	return log
}
func GetLogsById(id int64) Log {
	u := Log{LogId: id}
	o2 := orm.NewOrm()
	o2.Using("default")
	err := o2.Read(&u)
	if err == orm.ErrNoRows {
		log.Println("查询不到")
	} else if err == orm.ErrMissPK {
		log.Println("找不到主键")
	}
	return u
}

func DeleteLogs(id int64) {
	o := orm.NewOrm()
	o.Using("default")
	o.Delete(&Log{LogId: id})
}

//自动建表
func (this *Log) createTable() {
	name := "default"                          //数据库别名
	force := false                             //不强制建数据库
	verbose := true                            //打印建表过程
	err := orm.RunSyncdb(name, force, verbose) //建表
	if err != nil {
		//beego.Error(err)
	}
}


// 添加记录
func AddLogs(
	userid int64, //用户唯一id
	bookid int64, //书籍唯一id
	operateContent string, //操作内容
	operateDate time.Time,
) (int64, error) { //操作日期
	o := orm.NewOrm()
	o.Using("default")
	logs := new(Log)
	logs.UserId = userid //生成随机字符串
	logs.BookId = bookid
	logs.OperateContent = operateContent
	logs.OperateDate = operateDate
	logs.createTable()
	return o.Insert(logs)
}
