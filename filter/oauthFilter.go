package filter

import (
	"user/common"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var OAuthFilter = func(ctx *context.Context) {
	access_token := ctx.Input.Query("access_token")
	beego.Debug("OAuthFilter >>>>>>> access_token : ", access_token)
	if access_token == "" {
		beego.Debug("OAuthFilter >>>>>>> access_token is empty")
		ctx.Output.JSON(common.ResponseResult(common.ErrCodeNoPermission, nil), false, false)
		//ctx.WriteString("30414")
	}
	//封掉客户端 uid
	//redis 缓存中是否存在
	//通过token获取数据库用户信息
	//保持到redis缓存中

	//将数据放入到context中，方便后续controller处理
	ctx.Input.SetData("id", 1000002)
}
