package initialize

import (
	"fmt"
	"go_im/common/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func MysqlConnector() {
	m := global.Config.Mysql
	var dsn = fmt.Sprintf("%s:%s@%s", m.UserName, m.Password, m.Url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		fmt.Printf("mysql error: %s-----%s", err, dsn)
		return
	}
	sqlDb, err := db.DB()
	if err != nil {
		fmt.Printf("mysql error: %s", err)
	}

	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
	global.Db = db
}
