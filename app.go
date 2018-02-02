package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/xwinie/glue/lib/db"
	"github.com/xwinie/glue/router"
)

//Engine 获取engine
func main() {
	dbconfig := db.Config{}
	dbconfig.DbHost = "192.168.3.253"
	dbconfig.DbName = "yunhao"
	dbconfig.DbUser = "yunhao"
	dbconfig.DbPassword = "1qazZAQ!"
	dbconfig.DbPort = "3306"
	dbconfig.DbType = "mysql"
	dbconfig.DbCharset = "utf8mb4"
	dbconfig.DbPath = []string{"/Users/bobo/go/src/github.com/xwinie/glue/app"}

	s := gin.Default()
	err := db.Connect(dbconfig)
	if err != nil {
		log.Fatal("init db error:", err.Error())
	}

	router.Routers(s)
	s.Run()
}
