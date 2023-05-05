package db

import (
	"fmt"
	"goroutine/helpers/env"
)

var dbConfig *DbConfig = nil

func init() {
	if dbConfig == nil {
		defaultConfig := GetDefaultConfig()
		dbConfig = &defaultConfig
	}
}

type DbConfig struct {
	DbConnection string
	DbHost       string
	DbPort       string
	DbDatabase   string
	DbUsername   string
	DbPassword   string
}

func (config *DbConfig) GetDsn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DbUsername,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbDatabase,
	)
}

func GetDefaultConfig() DbConfig {
	if dbConfig == nil {
		dbConfig = &DbConfig{
			env.Get("DB_CONNECTION"),
			env.Get("DB_HOST"),
			env.Get("DB_PORT"),
			env.Get("DB_DATABASE"),
			env.Get("DB_USERNAME"),
			env.Get("DB_PASSWORD"),
		}
	}

	return *dbConfig // pass by value
}
