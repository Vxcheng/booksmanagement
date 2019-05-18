package services

import (
	"log"
	"ownergit/booksmanagement/models"
	"strconv"
	"time"

)

//用户借书 库存减少相应数量 生成日志时间操作状态
func BorrowB(uid, bid, num int) (str string, b bool) {
	book1, _ := models.GetOneBook(bid)
	if book1.Stock < 1 {
		log.Println("库存不足无法借书")
		str = "库存不足无法借书"
		b = false
	} else {
		stock := book1.Stock - num
		var b2 models.Book
		b2.BookId = book1.BookId
		b2.Stock = stock
		b2.BookName = book1.BookName
		models.UpdateOne(b2)
		t := time.Now().Local()
		opera := 1
		s := "用户id（"
		//uname := "徐静仪"    根据uid得到user取出uname
		str = s + strconv.Itoa(uid) + "）的用户" + models.OperateCount(opera) + "了" + strconv.Itoa(num) + "本" + b2.BookName
		AddLog(uid, bid, str, t)
		b = true
	}
	return
}

func ReturnB(uid, bid, num int) (str string, b bool) {
	book1, _ := models.GetOneBook(bid)
	stock := book1.Stock + num
	//var b2 models.Book
	//b2.BookId = bid
	//b2.Stock = stock
	//b2.BookName = book1.BookName
	b2 := models.Book{BookId: bid, BookName: book1.BookName, Stock: stock}
	models.UpdateOne(b2)
	t := time.Now().Local()
	opera := 0
	s := "用户id（"
	str = s + strconv.Itoa(uid) + "）的用户" + models.OperateCount(opera) + "了" + strconv.Itoa(num) + "本" + b2.BookName
	//新增一条日志
	AddLog(uid, bid, str, t)
	b = true
	return
}

//func AddLog(uid, bid int, op string, t time.Time) {
//	o := orm.NewOrm()
//	//var logs models.Logs
//	//logs.Uid = int64(uid)
//	//logs.Bookid = int64(bid)
//	//logs.OperateContent = op
//	//logs.OperateDate = t
//	logs := &models.Logs{Uid: int64(uid), Bookid: int64(bid), OperateContent: op, OperateDate: t}
//	o.Insert(logs)
//}

//add by csp
//add time:2018/7/14/13:59
func AddLog(uid int, bid int, op string, t time.Time) {
	models.AddLogs(int64(uid), int64(bid), op, t)
}

//add by csp endx
