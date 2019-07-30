package database

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

var Engine *xorm.Engine

var (
	db       = "vote_example"
	user     = "root"
	password = "admin123"
)

func MySQLDBInit() {
	dataSourceName := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=true&loc=Local", user, password, db)
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		log.Fatalln("CONNECT DATABASE FAIL :", err.Error())
		return
	}
	Engine = engine

	Engine.ShowSQL(true)
	Engine.SetMapper(core.GonicMapper{})
	Engine.SetTableMapper(core.GonicMapper{})
}
