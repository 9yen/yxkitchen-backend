package bootstrap

import (
	"errors"
	"fmt"
	"time"
	"yxkitchen-backend/conf"
	"yxkitchen-backend/pkg/database"
	"yxkitchen-backend/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {

	var dbConfig gorm.Dialector
	switch conf.C.Database.Connection {
	case "mysql":
		// 构建 DSN 信息
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			conf.C.Mysql.Username,
			conf.C.Mysql.Password,
			conf.C.Mysql.Host,
			conf.C.Mysql.Port,
			conf.C.Mysql.Database,
			conf.C.Mysql.Charset,
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	case "sqlite":
		// 初始化 sqlite
		database := conf.C.Sqlite.SqlFile
		dbConfig = sqlite.Open(database)
	default:
		panic(errors.New("database connection not supported"))
	}

	// 连接数据库，并设置 GORM 的日志模式
	database.Connect(dbConfig, logger.NewGormLogger())

	// 设置最大连接数
	database.SQLDB.SetMaxOpenConns(conf.C.Mysql.MaxOpenConnections)
	// 设置最大空闲连接数
	database.SQLDB.SetMaxIdleConns(conf.C.Mysql.MaxIdleConnections)
	// 设置每个链接的过期时间
	database.SQLDB.SetConnMaxLifetime(time.Duration(conf.C.Mysql.MaxLifeSeconds) * time.Second)
}
