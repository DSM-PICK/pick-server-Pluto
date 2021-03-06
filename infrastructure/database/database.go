package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

var DB *Database

func InitializeMysql() *Database {
	return initialize(mysqlConnection())
}

func initialize(dbConnection gorm.Dialector) *Database {
	if DB != nil { return DB }
	connection, e := gorm.Open(dbConnection, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if e != nil {
		log.Panic("database connection failed")
	}
	DB = &Database{Connection: connection}
	return DB
}

func mysqlConnection() gorm.Dialector {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		func() string {
			port := os.Getenv("PORT")
			if port == "" {
				port = "3306"
			}
			return port
		}(),
		"pick?charset=utf8&parseTime=True")
	return mysql.Open(url)
}

type Database struct {
	Connection *gorm.DB
}
