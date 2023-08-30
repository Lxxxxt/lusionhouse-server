package mysql

import (
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Nil      = gorm.ErrRecordNotFound
	MysqlCli *gorm.DB
)

func Startup() (err error) {
	dsn := "root:qusiba456@tcp(mysql-service:3306)/lusion_house"
	MysqlCli, err = initDb(dsn)
	return
}

func initDb(uri string) (*gorm.DB, error) {
	if !strings.HasSuffix(uri, "?parseTime=true") {
		uri = uri + "?parseTime=true"
	}
	return gorm.Open(mysql.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}
