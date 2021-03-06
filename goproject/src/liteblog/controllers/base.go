package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/satori/go.uuid"
	"liteblog/json"
	"liteblog/models"
	"liteblog/syserror"
)

//定义session中的key值
const SESSION_USER_KEY = "SESSION_USER_KEY"

type BaseController struct {
	beego.Controller
	IsLogin bool        //是否登录
	User    models.User //登录的有用户
}

type NestPreparer interface {
	NextPreparse()
}

//所有的路由都会继承此方法
//首先执行的方法
func (this *BaseController) Prepare() {

	//在匹配之前，判断用户是否登录，盘点session中是否存在用户，存在就已经登录，否则就没有登录

	this.IsLogin = false
	session := this.GetSession(SESSION_USER_KEY)
	logs.Info("session--->", session)
	if session != nil {
		//session是否为User对象
		if u, ok := session.(models.User); ok {
			this.User = u
			this.Data["user"] = u
			this.IsLogin = true
		}
	}

	logs.Info("是否登录过呢", this.IsLogin)

	this.Data["IsLogin"] = this.IsLogin

	//请求的url
	this.Data["path"] = this.Ctx.Request.RequestURI

	//如果实现了此方法后
	if preparer, ok := this.AppController.(NestPreparer); ok {
		preparer.NextPreparse()
	}

}

//判断是否非空
func (c *BaseController) GetMustString(key string, msg string) string {
	email := c.GetString(key, "")
	if len(email) == 0 {
		c.Abort500(errors.New(msg))
	}
	return email
}

//返回json时候的封装
func (ctx *BaseController) JSONOk(msg string, actions ...string) {
	var action string
	if len(actions) > 0 {
		action = actions[0]
	}
	ctx.Data["json"] = &json.ResultJsonValue{
		Code:   0,
		Msg:    msg,
		Action: action,
	}
	ctx.ServeJSON()
}

func (this *BaseController) Abort500(err error) {
	//请求的url
	this.Data["error"] = err
	this.Abort("500")
}

//是否登录过
func (ctx *BaseController) MustLogin() {
	if !ctx.IsLogin {
		ctx.Abort500(syserror.NoUserError{})
	}
}

func (this *BaseController) UUID() string {
	uuids, e := uuid.NewV4()
	if e != nil {
		this.Abort500(syserror.NewError("系统错误", e))
	}
	return uuids.String()
}
