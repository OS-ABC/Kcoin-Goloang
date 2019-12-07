package controllers

import (
	"Kcoin-Golang/src/models"
	_ "Kcoin-Golang/src/models"

	"github.com/astaxie/beego"
)

type PersonalPageController struct {
	beego.Controller
}

func (c *PersonalPageController) Get() {
	//jsonBuf是一个用于调试的静态json，之后会调用webserver的接口，动态获取。
	//jsonBuf :=
	//	`{
	//	"errorCode": "0",
	//	"data":{
	//		"userName": "DoubleJ",
	//		"headShotUrl": "../static/img/tx2.png"
	//	}
	//}`
	status := c.Ctx.GetCookie("status")
	c.Ctx.SetCookie("lastUri", c.Ctx.Request.RequestURI)
	if status == "0" || status == "" {
		defer c.Redirect("/login.html", 302)
	}
	user := models.UserInfo{Data: &models.UserData{}} //user中存放着json解析后获得的数据。
	user.Data.UserName = c.Ctx.GetCookie("userName")
	user.Data.HeadShotUrl = c.Ctx.GetCookie("headShotUrl")
	c.Data["user"] = user
	c.TplName = "personalPage.html" //该controller对应的页面
}
