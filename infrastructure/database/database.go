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

type ConnectionInformation struct {
	Username     string
	Password     string
	Host         string
	Port         string
	DatabaseName string
}

func DefaultInitialize() *Database {
	return Initialize(MysqlConnection())
}

func Initialize(dbConnection gorm.Dialector) *Database {
	if DB != nil {
		return DB
	}
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

func MysqlConnection() gorm.Dialector {
	url := MysqlUrl(ConnectionInformation{
		Username: os.Getenv("DATABASE_USERNAME"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Host: os.Getenv("DATABASE_HOST"),
		Port: os.Getenv("PORT"),
		DatabaseName: "pick",
	})
	return mysql.Open(url)
}

func MysqlUrl(information ConnectionInformation) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		information.Username,
		information.Password,
		information.Host,
		func() string {
			port := information.Port
			if port == "" {
				port = "3306"
			}
			return port
		}(),
		information.DatabaseName+"?charset=utf8&parseTime=True")
}

type Database struct {
	Connection *gorm.DB
}
