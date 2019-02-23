package models

import (
	"log"

	"github.com/astaxie/beego/orm"
)

type Book struct {
	BookId   int    `json:"bookId" orm:"pk;auto;column(bookId)" `
	BookName string `json:"bookName" orm:"null;column(bookName)" `
	Stock    int    `json:"stock" orm:"null;column(stock)" `
}

func AddBook(books Book) int {
	orm.NewOrm().Insert(&books)
	log.Println(books)
	return books.BookId
}

func GetAllBook(page, pageSize int) (int, []*Book) {
	offset := pageSize * (page - 1)
	bookList := make([]*Book, pageSize)
	book := new(Book)
	o := orm.NewOrm().QueryTable(book)
	total, _ := o.Count()
	log.Println("GetAllBook", offset, total)
	o.Limit(pageSize, offset).OrderBy("bookId").All(&bookList)

	log.Println("booksList", bookList)
	return int(total), bookList
}

func UpdateOne(bo Book) Book {
	orm.NewOrm().Update(&bo)
	//book := new(Book)
	//err := orm.NewOrm().Raw("SELECT * FROM book WHERE BookId = ?", bookid).QueryRow(&book)
	//fmt.Println("bookid_update:", bo)
	//if err != nil {
	//	log.Println("UpdateOne_err:", err)
	//} else {
	//	orm.NewOrm().QueryTable(book).Filter("BookId", bookid).Update(orm.Params{
	//		"BookName": bo.BookName, "Stock": bo.Stock})
	//	//orm.NewOrm().Raw("UPDATE book SET bookName = ?,stock = ? ",bo.BookName,bo.BookName).Values(&b)
	//}
	//return GetOne(bookid)
	return bo
}

func GetOneBook(bookid int) (*Book, error) {
	//book := new(Books)
	b := &Book{BookId: bookid}
	err := orm.NewOrm().Read(b)
	return b, err
}

func DeleteOne(id int) error {
	book := &Book{BookId: id}
	_, err := orm.NewOrm().Delete(book)
	return err
}

func OperateCount(i int) string {
	var s string
	if i == 0 {
		s = "还"
	} else {
		s = "借"
	}
	return s
}
