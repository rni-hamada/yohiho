package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Handler *gorm.DB

func Connect() *gorm.DB {
	user := getEnvOrDefault("MYSQL_USER", "root")
	pass := getEnvOrDefault("MYSQL_PASSWORD", "")
	host := getEnvOrDefault("MYSQL_HOST", "localhost")
	port := getEnvOrDefault("MYSQL_PORT", "3306")
	dbname := getEnvOrDefault("MYSQL_DBNAME", "yodb")
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, dbname))
	if err != nil {
		panic("ohno")
	}
	db.LogMode(true)
	Handler = db
	return db
}

func getEnvOrDefault(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
