package controllers

import (
	"net/http"
	"strings"
	"time"
	"user/common"
	"user/models"
	"user/util"

	"github.com/astaxie/beego"
)

type ThirdloginController struct {
	beego.Controller
}

const (
	//H5_APP_ID     = "wx87f81569b7e4b5f6"
	//H5_APP_SECRET = "8421fd4781b1c29077c2e82e71ce3d2a"
	WEIXIN_URL = "https://api.weixin.qq.com/sns/"
)

var (
	H5_APP_ID     string = beego.AppConfig.String("h5.app.id")
	H5_APP_SECRET string = beego.AppConfig.String("h5.app.secret")
)

//test url : https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx87f81569b7e4b5f6&redirect_uri=http%3a%2f%2ftest-user.lezhuale.com&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect
// @router /weixin_h5 [get]
func (u *ThirdloginController) Weixin_h5() {
	code := u.Input().Get("code")
	access_token := u.Input().Get("access_token")
	redirect_url := u.Input().Get("url")

	//获取用户accesstoken和openid
	token_url := WEIXIN_URL + "oauth2/access_token?grant_type=authorization_code&appid=" + H5_APP_ID + "&secret=" + H5_APP_SECRET + "&code=" + code
	resp, err := util.GetJSON(token_url, 5*time.Second)
	if err != nil {
		beego.Error("weixin_h5 error :", err)
	}
	if _, exists := resp["access_token"]; !exists {
		u.Data["json"] = common.ResponseResult(common.ErrCodeInvalidParams, nil)
		u.ServeJSON()
		return
	}
	access_token = resp["access_token"].(string)
	openid := resp["openid"].(string)
	beego.Debug("token", access_token, "openid:", openid)

	//获取用户信息 unionid等等
	userInfo_url := WEIXIN_URL + "userinfo?access_token=" + access_token + "&openid=" + openid
	resp, err = util.GetJSON(userInfo_url, 5*time.Second)
	if err != nil {
		beego.Error("weixin_h5 get userInfo error :", err)
	}
	beego.Debug("weixin_h5 userInfo :", resp)
	unionid := resp["unionid"].(string)
	nickname := resp["nickname"].(string)
	pic := resp["headimgurl"].(string)
	sex := resp["sex"].(float64) //1:男, 0:女
	//用户是否已经存在
	user := models.GetUser("weixin_unionid1", unionid)
	//生成新用户
	if user.Id_ == 0 {
		newUser := &models.User{
			Weixin_openid:       openid,
			Weixin_access_token: access_token,
			Weixin_unionid:      unionid,
			Pic:                 pic,
			Nickname:            nickname,
			Via:                 "weixin",
			Sex:                 sex,
		}
		beego.Debug("weixin_h5 user :", &user)
		user, err = models.BuildUser(newUser)
		if err != nil {
			u.Data["json"] = common.ResponseResult(common.ErrCodeDuplicated, "")
			u.ServeJSON()
			return
		}
	}
	beego.Debug("redirect_url:", redirect_url)
	//转跳地址
	if redirect_url != "" {
		redirect_url = getRedirectUrlWithToken(redirect_url, user.Token)
		u.Redirect(redirect_url, http.StatusFound)
		return
	}
	u.Data["json"] = common.ResponseResult(common.ErrCodeOk, user)
	u.ServeJSON()
}

func getRedirectUrlWithToken(originUrl, token string) string {
	if strings.Contains(originUrl, "?") {
		originUrl = originUrl + "&"
	} else {
		originUrl = originUrl + "?"
	}
	if strings.Contains(originUrl, "{access_token}") {
		originUrl = strings.Replace(originUrl, "{access_token}", token, 1)
	} else {
		originUrl = originUrl + "access_token=" + token
	}
	beego.Debug("replace url : ", originUrl)
	return originUrl
}
