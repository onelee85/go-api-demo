package db

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/globalsign/mgo"
)

var globalMgoSession *mgo.Session

func init() {
	mongoHost := beego.AppConfig.String("mongo.host")
	beego.Informational("mongoHost:", mongoHost)
	session, err := mgo.DialWithTimeout(mongoHost, 30*time.Second)
	if err != nil {
		beego.Error("init mongo db", err)
	}
	globalMgoSession = session
	globalMgoSession.SetMode(mgo.Monotonic, true)
	globalMgoSession.SetPoolLimit(100)
}

/**
example:
	session := CloneSession() //传入数据库的地址，可以传入多个，具体请看接口文档
	defer session.Close() //用完记得关闭
	c := session.DB("test").C("people")
**/
func MongoSession() *mgo.Session {
	return globalMgoSession.Clone()
}

type MongoDataBase struct {
	dbName string
}

//数据库注册
var UserDb = &MongoDataBase{dbName: "account"}

//对集合执行操作 增删查改, 保证session能被close
func (db *MongoDataBase) Execute(collection string, do func(*mgo.Collection) error) error {
	session := MongoSession()
	defer func() {
		session.Close()
		beego.Debug("close session...")
		if err := recover(); err != nil {
			beego.Error("Execute mongo ", err)
		}
	}()
	c := session.DB(db.dbName).C(collection)
	return do(c)
}
