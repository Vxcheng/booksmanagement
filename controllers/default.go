//add by xjy

package controllers

import (
	"booksmanagementSys/models"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// Index
// @router /index [Get]
func (c *UserController) Get() {
	c.TplName = "index.tpl"
}

/****************************************************************
【函数】：GetMd5String()
【功能】:password MD5加密
【传入参数】:str
【输出参数】:string
【创建时间】：2018.07.06 by xjy
*******************************************************************/
func (c *UserController) GetMd5String(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	a := hex.EncodeToString(md5Ctx.Sum(nil))
	return a
}

type RegisterReq struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

/****************************************************************
【函数】：Register()
【功能】: 用户根据username和password注册
【传入参数】:
【输出参数】:
【创建时间】：2018.07.06 by xjy
*******************************************************************/
// Register
// @router /register [post]
func (c *UserController) Register() {
	var v RegisterReq

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		md5pwd := c.GetMd5String(v.Password)
		var u models.User
		u.Uname = v.Username
		u.Psw = md5pwd
		u.CreateTime = time.Now().UTC()
		u.LastTime = time.Now().UTC()
		if _, err := models.AddUser(&u); err == nil {
			c.Data["json"] = u
		} else {
			fmt.Println(err.Error())
			c.Data["json"] = err.Error()
		}
	} else {
		fmt.Println(err.Error())
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

type LoginReq struct {
	LoginId  int64  `json:"LoginId"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

/****************************************************************
【函数】：Login()
【功能】: 用户登陆
【传入参数】:
【输出参数】:
【创建时间】：2018.07.06 by xjy
*******************************************************************/
// Login
// @router /login [post]
func (c *UserController) Login() {
	var v LoginReq
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		md5pwd := c.GetMd5String(v.Password)
		usrOld := &models.User{Id: v.LoginId, Uname: v.Username, Psw: md5pwd}
		user, _ := models.VerifyUser(usrOld)
		if user != nil {
			// get uid

			id := user[0]["Id"]
			fmt.Println(id)
			c.Data["json"] = map[string]interface{}{"success": 0, "msg": "登录成功", "uname": v.LoginId, "id": id}
		} else {
			c.Data["json"] = map[string]interface{}{"success": -1, "msg": "账号密码错误"}
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
