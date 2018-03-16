package models

import (
	"strconv"
	"user/db"
	"user/util"

	"github.com/astaxie/beego"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type User struct {
	Id_                 int64   `bson:"_id" json:"_id"`
	Username            string  `json:"username"	bson:"username"`
	Timestamp           int64   `json:"timestamp"	bson:"timestamp"`
	Pwd                 string  `json:"-"`
	Pic                 string  `json:"pic"`
	Nickname            string  `json:"nickname"`
	Via                 string  `json:"via"`
	Sex                 float64 `json:"sex"`
	Token               string  `json:"token"`
	Weixin_openid       string  `json:"-"`
	Weixin_access_token string  `json:"-"`
	Weixin_unionid      string  `json:"-"`
}

func GetUser(field, val string) (user *User) {
	user = &User{}
	query := bson.M{}
	query[field] = val
	db.UserDb.Execute("users", func(c *mgo.Collection) {
		err := c.Find(query).One(&user)
		if err != nil {
			beego.Error("GetUser error:", err)
		}

	})
	return user
}

func ExistsUser(field, val string) bool {
	query := bson.M{}
	query[field] = val
	var count int
	var err error
	db.UserDb.Execute("users", func(c *mgo.Collection) {
		count, err = c.Find(query).Count()
		if err != nil {
			beego.Error("GetUser error:", err)
		}

	})
	return count > 0
}

//创建用户
func BuildUser(user *User) *User {
	user.Id_ = util.UserKGS.NextId()
	user.Timestamp = util.GetTimestampInMilli()
	user.Token = generateToken(user.Id_)
	db.UserDb.Execute("users", func(c *mgo.Collection) {
		err := c.Insert(user)
		if err != nil {
			beego.Error("bulid user error:", err)
		}
	})
	return user
}

func generateToken(uid int64) string {
	return util.Md5(strconv.FormatInt(uid, 36) + util.GetTimestampInMilliString() + util.RandString(16, nil))
}
