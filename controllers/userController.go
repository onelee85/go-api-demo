package controllers

import (
	"user/common"
	"user/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// @router /info [get]
func (u *UserController) Info() {
	access_token := u.Input().Get("access_token")
	//filter 内放置的id
	id := u.Ctx.Input.GetData("id").(int)
	beego.Debug("access_token : ", access_token, " id : ", id)
	user := models.GetUser("_id", id)
	u.Data["json"] = common.ResponseResult(common.ErrCodeOk, user)
	u.ServeJSON()
}
