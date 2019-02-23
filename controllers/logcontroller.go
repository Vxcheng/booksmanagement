//WRITE BY CSP
package controllers

import (
	"time"

	"github.com/astaxie/beego"
)
import (
	"encoding/json"
	"fmt"
	"booksmanagementSys/models"
)

type LogsController struct {
	beego.Controller
}

// @router / [get]
func (u *LogsController) GetAll() {
	fmt.Println("getLogs:")
	ss := models.GetAllLogs()
	ss2 := make([]models.Log4, len(ss))
	for n, v := range ss {
		ss2[n] = models.Log4{v.LogId, v.Username, v.Bookname, v.OperateContent, models.JsonTime(v.OperateDate)}
	}
	u.Data["json"] = ss2
	u.ServeJSON()
}

// @router /:id [get]
func (u *LogsController) GetById() {

	id, _ := u.GetInt64(":id")
	s := models.GetLogsById(id)
	u.Data["json"] = s
	u.ServeJSON()
}

// @router /:id [delete]
func (u *LogsController) Delete() {

	id, _ := u.GetInt64(":id")
	models.DeleteLogs(id)
	u.Data["json"] = true
	u.ServeJSON()
}

// @router / [post]
func (this *LogsController) Addlog() {
	var s models.Log
	json.Unmarshal(this.Ctx.Input.RequestBody, &s)
	userId := s.UserId
	bookId := s.BookId
	operation := s.OperateContent
	t := time.Now()
	id, err := models.AddLogs(userId, bookId, operation, t)
	if nil != err {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "add failed", "refer": "/"}
	} else {
		this.Data["json"] = map[string]interface{}{"result": true, "msg": fmt.Sprintf("[%d] ", id) + "add success", "refer": "/"}
	}
	this.ServeJSON()
}
