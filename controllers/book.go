package controllers

import (
	"encoding/json"
	"fmt"
	"booksmanagementSys/models"
	"booksmanagementSys/services"


	"github.com/astaxie/beego"
)

type BooksController struct {
	beego.Controller
}

type ResResult struct {
	ResCode int         `json:"resCode"`
	ResMsg  string      `json:"resMsg"`
	ResData interface{} `json:"resData"`
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (b *BooksController) Post() {
	//fmt.Println("test")
	var book models.Book
	json.Unmarshal(b.Ctx.Input.RequestBody, &book)
	id := models.AddBook(book)
	maps := make(map[string]interface{})
	maps["bookId"] = id
	res := &ResResult{0, "AddBook Success", maps}
	b.Data["json"] = res
	b.ServeJSON()
}

// @Title GetAllUser
// @Description get alluser
// @Param	 body	 page 	int 	true 	"fenye filter"
// @Param	 body	 pageSize 	int 	true 	"fenye filter"
// @Success 200 {Object} []models.Book
// @Failure 403 body is empty
// @router / [get]
func (b *BooksController) GetAll() {
	page, _ := b.GetInt("page")
	pageSize, _ := b.GetInt("pageSize")
	total, booksList := models.GetAllBook(page, pageSize)
	fmt.Println("GetAll:", booksList, total)
	maps := make(map[string]interface{})
	maps["total"] = total
	maps["data"] = booksList
	res := &ResResult{0, "GetAllBook  Success", maps}
	b.Data["json"] = res
	b.ServeJSON()
}

// @Title GetOne
// @Description getone user
// @Param	bookid		path 	string	true		"The uid you want to selelt onebook"
// @Success 200 {object} models.Book
// @Failure 403 :bookid is not int
// @router /:bookid [get]
func (b *BooksController) GetOne() {
	bookid, _ := b.GetInt(":bookid")
	fmt.Println("bookid:", bookid)
	book, err := models.GetOneBook(bookid)
	if err != nil {
		res := &ResResult{1, "GetOneBook Defailed", nil}
		b.Data["json"] = res
	} else {
		res := &ResResult{0, "GetOneBook Success", book}
		b.Data["json"] = res
	}
	b.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	bookid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :bookid is not int
// @router /:bookid [put]
func (b *BooksController) Put() {
	var book models.Book
	bookid, _ := b.GetInt(":bookid")
	book.BookId = bookid
	json.Unmarshal(b.Ctx.Input.RequestBody, &book)
	bo := models.UpdateOne(book)
	res := &ResResult{0, "UpdateOne Success", bo}
	b.Data["json"] = res
	b.ServeJSON()
}

// @router /:bookid [delete]
func (b *BooksController) Delete() {
	bookid, _ := b.GetInt(":bookid")
	fmt.Println("bookid:", bookid)
	err := models.DeleteOne(bookid)
	if err != nil {
		fmt.Println("delete_err:", err)
	}
	maps := make(map[string]interface{})
	maps["affectedId"] = bookid
	res := &ResResult{0, "DeleteOne  Success", maps}
	b.Data["json"] = res
	b.ServeJSON()
}

// @Title borrowbook
// @Description  user borrowbook generate log
// @Param	bookid		path 	string	true		"The uid you want to delete"
// @Success 200 {int} bookid
// @Failure 403 :bookid is not int
// @router /borrowbook [post]
func (b *BooksController) BorrowBook() {
	uid, _ := b.GetInt("uid")
	bid, _ := b.GetInt("bid")
	num, _ := b.GetInt("num")
	action, bools := services.BorrowB(uid, bid, num)
	maps := make(map[string]interface{})
	maps["operate"] = action
	if !bools {
		res := &ResResult{1, "BorrowBook Defailed", maps}
		b.Data["json"] = res
	} else {
		res := &ResResult{0, "BorrowBook Success", maps}
		b.Data["json"] = res
	}
	b.ServeJSON()
}

//
// @router /returnbook [post]
func (b *BooksController) ReturnBook() {
	uid, _ := b.GetInt("uid")
	bid, _ := b.GetInt("bid")
	num, _ := b.GetInt("num")
	action, bools := services.ReturnB(uid, bid, num)
	maps := make(map[string]interface{})
	maps["operate"] = action
	if !bools {
		res := &ResResult{1, "ReturnBook Defailed", maps}
		b.Data["json"] = res
	} else {
		res := &ResResult{0, "ReturnBook Success", maps}
		b.Data["json"] = res
	}
	b.ServeJSON()
}
