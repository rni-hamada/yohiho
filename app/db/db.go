package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rni-hamada/yohiho/app/util"
)

var Handler *gorm.DB

func Connect(user, pass, host, port, dbname string) *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, dbname))
	if err != nil {
		panic("ohno")
	}
	return db
}

func Prepare() {
	user := util.GetEnvOrDefault("MYSQL_USER", "root")
	pass := util.GetEnvOrDefault("MYSQL_PASSWORD", "")
	host := util.GetEnvOrDefault("MYSQL_HOST", "localhost")
	port := util.GetEnvOrDefault("MYSQL_PORT", "3306")
	dbname := util.GetEnvOrDefault("MYSQL_DBNAME", "yodb")
	Handler = Connect(user, pass, host, port, dbname)
	Handler.LogMode(true)
}
