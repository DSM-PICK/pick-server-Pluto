package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

var DB *Database

func Initialize() *Database {
	if DB != nil { return DB }
	connection, e := gorm.Open(mysql.Open(databaseUrl()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if e != nil {
		log.Panic("database connection failed")
	}
	DB = &Database{connection: connection}
	return DB
}

func databaseUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
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
}

type Database struct {
	connection *gorm.DB
}
