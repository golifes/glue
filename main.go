package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/xwinie/glue/core"
	"github.com/xwinie/glue/migrate"
	"github.com/xwinie/glue/router"
)

func main() {
	dbconfig := core.Config{}
	dbconfig.DbType = "mysql"
	dbconfig.DbHost = "192.168.3.253"
	dbconfig.DbName = "yunhao"
	dbconfig.DbUser = "yunhao"
	dbconfig.DbPort = "3306"
	dbconfig.DbPassword = "1qazZAQ!"
	dbconfig.DbCharset = "utf8mb4"
	dbconfig.DbShowSQL = true
	dbconfig.DbPath = []string{"/Users/bobo/go/src/github.com/xwinie/glue/app"}
	err := core.Connect(dbconfig)
	if err != nil {
		log.Fatal("init db error:", err.Error())
	}

	o := core.New()
	migrate.Migrate(o)
	s := echo.New()
	// Middleware
	s.Use(middleware.Logger())
	// s.Use(middleware.Recover())
	router.Routers(s)
	s.Logger.Fatal(s.Start(":1323"))
}
