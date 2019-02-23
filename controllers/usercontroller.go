//add by Chengshipeng
package controllers

import (
	_ "time"

	"booksmanagementSys/models"

	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gogather/com"
)

// 注册
type RegistorController struct {
	beego.Controller
}

// @router / [get]
func (this *RegistorController) Get() {
	registorable, err := beego.AppConfig.Bool("registorable")
	if registorable || nil != err {
		//this.TplName = "registor.tpl"
	} else {
		this.Ctx.WriteString("registor closed")
	}
}

// @router / [post]
func (this *RegistorController) Post() {
	//registorable, err := beego.AppConfig.Bool("registorable")
	//	if nil != err {
	//		// default registorable is true, do nothing
	//	} else if !registorable {
	//		this.Data["json"] = map[string]interface{}{"result": false, "msg": "registorable is false", "refer": "/"}
	//		this.ServeJSON()
	//		return
	//	}
	//username := this.GetString("username")
	//username := this.Ctx.Input.RequestBody("username")
	var s models.LibUser
	json.Unmarshal(this.Ctx.Input.RequestBody, &s)
	//username := s.GetString("username")
	username := s.Username
	if username == "" {
		//this.Ctx.WriteString("usernaem is empty")
		//return
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "illegal username", "refer": "/"}
		this.ServeJSON()
		return
	}
	//log.Println(username)
	//password := this.GetString("password")
	password := s.Password
	if password == "" {
		//this.Ctx.WriteString("usernaem is empty")
		//return
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "illegal username", "refer": "/"}
		this.ServeJSON()
		return
	}
	/*if !utils.CheckUsername(username) {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "illegal username", "refer": "/"}
		this.ServeJSON()
		return
	}*/
	_, err := models.AddUsers(username, password)
	if nil != err {
		this.Data["json"] = map[string]interface{}{"result": false}
	} else {
		this.Data["json"] = map[string]interface{}{"result": true}
	}
	this.ServeJSON()
}

// 登录
type LoginController struct {
	beego.Controller
}

//// @router / [get]
//func (this *LoginController) Get() {
//	// if not login, permission deny
//	user := this.GetSession("username")
//	if user != nil {
//		//this.Redirect("/admin", 302)
//	} else {
//		//this.TplName = "login.tpl"
//	}
//}

// @router / [post]
func (this *LoginController) Post() {
	//username := this.GetString("username")
	//password := this.GetString("password")
	var s models.LibUser
	json.Unmarshal(this.Ctx.Input.RequestBody, &s)
	username := s.Username
	password := s.Password
	if username == "" || password == "" {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	}
	user, err := models.FindUser(username)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "user does not exist", "refer": "/"}
	} else {
		passwd := com.Md5(password + user.Salt)
		//log.Println(password)
		//log.Println(passwd)
		if passwd == user.Password {
			//this.SetSession("username", username)
			this.Data["json"] = map[string]interface{}{"result": true, "userId": user.Id}
		} else {
			this.Data["json"] = map[string]interface{}{"result": false, "userId": "Null"}
		}
	}
	this.ServeJSON()
}

// 登出
type LogoutController struct {
	beego.Controller
}

//// @router / [get]
//func (this *LogoutController) Get() {
//	this.DelSession("username")
//	this.Ctx.WriteString("you have logout")
//}

// @router /:id [post]
func (this *LogoutController) Post() {
	//_, _ := this.GetInt64(":id")
	this.Data["json"] = map[string]interface{}{"result": true}
	//this.Ctx.WriteString("you have logout")
	this.ServeJSON()
}

//// 修改用户名
//type ChangeUsernameController struct {
//	beego.Controller
//}

//// @router / [get]
//func (this *ChangeUsernameController) Get() {
//	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request ", "refer": "/"}
//	this.ServeJSON()
//}

//// @router / [post]
//func (this *ChangeUsernameController) Post() {
//	// if not login, permission deny
//	user := this.GetSession("username")
//	if user == nil {
//		this.Data["json"] = map[string]interface{}{"result": false, "msg": "login first please", "refer": nil}
//		this.ServeJSON()
//		return
//	}

//	oldUsername := user.(string)
//	newUsername := this.GetString("username")

//	err := models.ChangeUsername(oldUsername, newUsername)

//	if nil != err {
//		// log.Println(err)
//		this.Data["json"] = map[string]interface{}{"result": false, "msg": "change username failed", "refer": "/"}
//		this.ServeJSON()
//	} else {
//		this.SetSession("username", newUsername)
//		this.Data["json"] = map[string]interface{}{"result": true, "msg": "change username success", "refer": "/"}
//		this.ServeJSON()
//	}
//}

//// 设置密码
//type SetPasswordController struct {
//	beego.Controller
//}

//// @router /:varify [get]
//func (this *SetPasswordController) Get() {
//	varify := this.Ctx.Input.Param(":varify")

//	if "" == varify {
//		this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request ", "refer": "/"}
//		this.ServeJSON()
//	}

//	result, username, err := CheckVarify(varify)

//	if nil != err {
//		this.Ctx.WriteString("找回密码已过期")
//	} else if !result {
//		this.Ctx.WriteString("验证错误")
//	} else {
//		this.Data["username"] = username
//		this.SetSession("username", username)
//		this.SetSession("reset", true)
//		//this.TplName = "resetpasswd.tpl"
//	}
//}

//// @router /:varify [post]
//func (this *SetPasswordController) Post() {
//	user := this.GetSession("username")
//	reset := this.GetSession("reset")
//	resetable := reset.(bool)
//	if !resetable {
//		this.Data["json"] = map[string]interface{}{"result": false, "msg": "resetable is false", "refer": nil}
//		this.ServeJSON()
//		return
//	}
//	if user == nil {
//		this.Data["json"] = map[string]interface{}{"result": false, "msg": "session failed", "refer": nil}
//		this.ServeJSON()
//		return
//	}
//	username := user.(string)
//	newPassword := this.GetString("password")
//	// fmt.Println(username)
//	if "" == newPassword {
//		this.Data["json"] = map[string]interface{}{"result": false, "msg": "password is needed", "refer": nil}
//		this.ServeJSON()
//		return
//	}

//	err := SetPassword(username, newPassword)
//	if nil != err {
//		this.Data["json"] = map[string]interface{}{"result": false, "msg": "set password failed", "refer": nil}
//		this.ServeJSON()
//		return
//	} else {
//		this.DelSession("reset")
//		this.Data["json"] = map[string]interface{}{"result": true, "msg": "set password success", "refer": "/"}
//		this.ServeJSON()
//	}

//}

//// 修改密码
//type ChangePasswordController struct {
//	beego.Controller
//}

//// @router / [get]
//func (this *ChangePasswordController) Get() {
//	this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request ", "refer": "/"}
//	this.ServeJSON()
//}

//// @router / [post]
//func (this *ChangePasswordController) Post() {
//	// if not login, permission deny
//	user := this.GetSession("username")
//	if user == nil {
//		this.Data["json"] = map[string]interface{}{"result": false, "msg": "login first please", "refer": nil}
//		this.ServeJSON()
//		return
//	}

//	username := user.(string)
//	oldPassword := this.GetString("old_password")
//	newPassword := this.GetString("password")

//	err := ChangePassword(username, oldPassword, newPassword)

//	if nil != err {
//		this.Data["json"] = map[string]interface{}{"result": false, "msg": "change password faild", "refer": nil}
//		this.ServeJSON()
//	} else {
//		this.Data["json"] = map[string]interface{}{"result": true, "msg": "change password success", "refer": nil}
//		this.ServeJSON()
//	}
//}
