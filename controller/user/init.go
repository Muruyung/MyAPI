package user

import "xorm.io/xorm"

type Database struct {
	engine *xorm.Engine
}

func Init(DB *xorm.Engine) Database {
	return Database{engine: DB}
}
