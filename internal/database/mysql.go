package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simple-blog-api-golang/configs"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}

	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		configs.Env.DbUser,
		configs.Env.DbPassword,
		configs.Env.DbHost,
		configs.Env.DbPort,
		configs.Env.DbName)

	var err error

	db, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
