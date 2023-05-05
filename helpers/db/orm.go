package db

import (
	"database/sql"
	"goroutine/helpers/env"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	DB = NewConnection(dbConfig)
}

// NewConnection Get a new database connection
func NewConnection(config *DbConfig) *gorm.DB {

	dsn := config.GetDsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connection to mysql failed:" + err.Error())

	}

	// print each queries if GORM debug mode is enabled
	// see: https://gorm.io/zh_CN/docs/logger.html
	var gormDebugMode string = env.Get("GORM_DEBUG")
	enabled, err := strconv.ParseBool(gormDebugMode)
	if (err == nil) && enabled {
		db.Config.Logger = logger.Default.LogMode(logger.Info)
	} else {
		db.Config.Logger = logger.Default.LogMode(logger.Error)
	}

	return db
}

func CloseConnection() {
	conn, _ := DB.DB()
	conn.Close()
}

func BeginTransaction() {
	DB = NewConnection(dbConfig).Begin()
}

func Rollback() {
	tx := DB.Rollback()
	if err := tx.Error; err != nil && err != sql.ErrTxDone {
		tx.AddError(err)
	}
}

func Commit() {
	DB.Commit()
}
