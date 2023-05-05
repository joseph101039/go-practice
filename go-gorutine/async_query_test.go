package main

import (
	"fmt"
	"goroutine/helpers/env"
	"goroutine/helpers/goerror"
	"goroutine/models"
	"log"
	"sync"
	"testing"

	_ "goroutine/helpers/env" // loadenv

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var googleUser models.GoogleUser
var OauthToken models.OauthToken

func Test_syncQuery(t *testing.T) {

	syncQuery(t)
}

func Benchmark_syncQuery(b *testing.B) {

	syncQuery(b)
}

func Test_asyncQuery(t *testing.T) {
	asyncQuery(t)
}

func syncQuery(tb interface{}) {

	db := newConn()
	syncSampleQuery(db)
}

func asyncQuery(tb interface{}) {
	db := newConn()

	ch := make(chan *models.OauthToken, 5)
	const pcount = 10
	var wg sync.WaitGroup
	for i := 0; i < pcount; i++ {
		go asyncSampleQuery(db, ch)
		wg.Add(1)
	}

	var (
		count uint = 0
	)

	for result := range ch {
		count++

		log.Printf("count %d,\n %s\n", count, result)
		if count == pcount {
			close(ch)
			break
		}
	}
}

func asyncSampleQuery(db *gorm.DB, ch chan<- *models.OauthToken) {

	db.Begin()
	defer db.Rollback()

	result := new(models.OauthToken)
	err := db.Model(OauthToken).
		Joins(fmt.Sprintf("JOIN %s ON %s.id = %s.google_user_id",
			googleUser.TableName(),
			googleUser.TableName(),
			OauthToken.TableName(),
		)).
		Where("1=1").
		Scan(result).Error

	goerror.Fatal(err)
	db.Commit()

	ch <- result
}

func syncSampleQuery(db *gorm.DB) {
	db.Begin()
	defer db.Rollback()

	result := new(models.OauthToken)
	err := db.Model(OauthToken).
		Joins(fmt.Sprintf("JOIN %s ON %s.id = %s.google_user_id",
			googleUser.TableName(),
			googleUser.TableName(),
			OauthToken.TableName(),
		)).
		Where("1=1").
		Scan(result).Error

	goerror.Fatal(err)
	db.Commit()

	log.Printf("%#v\n", result)
	log.Println(result)
}

func newConn() *gorm.DB {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",

		env.Get("DB_USERNAME"),
		env.Get("DB_PASSWORD"),
		env.Get("DB_HOST"),
		env.Get("DB_PORT"),
		env.Get("DB_DATABASE"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // global disable to gain 30%+ performance improvement
		AllowGlobalUpdate:      true,
		Logger:                 logger.Default.LogMode(logger.Error), // print every queries
	})
	goerror.Fatal(err)

	sqlDB, err := db.DB()
	goerror.Fatal(err)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db

}
