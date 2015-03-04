package db

import (
	"github.com/go-xorm/xorm"
	"log"
)

const (
	DB_HOST = "127.0.0.1:3306"
	DB_USER = "root"
	DB_PWD = "test"
	DB_NAME = "test"
)
func GetEngine ()(*xorm.Engine){
	engine, err := xorm.NewEngine("mysql", DB_USER+":"+DB_PWD+"@/"+DB_NAME+"?charset=utf8")
	if err!=nil{
		log.Println(err)
	}
	engine.ShowSQL=true
	return engine
}
