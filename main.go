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
	dbHost     = os.Getenv("DB_CONTAINER_NAME")
	appPort    = os.Getenv("APP_PORT")
	dbPort     = os.Getenv("DB_PORT")
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbDatabase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("fail to connect DB")
		return
	}
	port := fmt.Sprintf(":%s", appPort)
	router := api.NewRouter(db)

	if err := router.Run(port); err != nil {
		panic(err)
	}

}
