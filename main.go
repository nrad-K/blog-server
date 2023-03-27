package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nrad-K/blog-server/api"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbUser     = os.Getenv("USERNAME")
	dbPass     = os.Getenv("USERPASSWD")
	dbDatabase = os.Getenv("DATABASE")
	dbHost     = os.Getenv("HOSTNAME")
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbDatabase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("fail to connect DB")
		return
	}

	router := api.NewRouter(db)

	if err := router.Run(":8000"); err != nil {
		panic(err)
	}

}
