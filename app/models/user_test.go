package models

import (
	"os"
	"testing"

	"github.com/rni-hamada/yohiho/app/db"
	"github.com/rni-hamada/yohiho/app/util"
	"github.com/stretchr/testify/assert"
)

func connectTestDb() {
	user := util.GetEnvOrDefault("MYSQL_USER", "root")
	pass := util.GetEnvOrDefault("MYSQL_PASSWORD", "")
	host := util.GetEnvOrDefault("MYSQL_HOST", "localhost")
	port := util.GetEnvOrDefault("MYSQL_PORT", "3306")
	dbname := util.GetEnvOrDefault("MYSQL_DBNAME", "yodb")
	db.Handler = db.Connect(user, pass, host, port, dbname)
}

func TestMain(m *testing.M) {
	connectTestDb()

	code := m.Run()

	os.Exit(code)
}

func TestGetUser(t *testing.T) {
	u := GetUser()
	assert.Equal(t, uint(1), u.ID)
}
