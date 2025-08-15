package db

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pole-audit/pkg/config"
	"time"
)

var Instance *gorm.DB

func init() {
	var err error
	dsn := config.Properties.DSN + "/test?charset=utf8&parseTime=True&loc=Local"
	if Instance, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		panic(err)
	}

	var sqlDB *sql.DB
	if sqlDB, err = Instance.DB(); err != nil || sqlDB.Ping() != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Hour * 24)
}
