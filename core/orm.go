package core

import (
	"fmt"
	"net/url"

	"github.com/xormplus/xorm"
)

// Config database orm config
type Config struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbName     string
	DbType     string
	DbPort     string
	DbPath     []string
	DbCharset  string
	DbShowSql  bool
}

//engine 定义全局变量
var engine *xorm.Engine

// Connect connect database return osm struct
func Connect(config Config) (err error) {
	if config.DbType == "sqlite3" {
		engine, err = xorm.NewEngine("sqlite3", "file::memory:?mode=memory&cache=shared&loc=Local&parseTime=true")
	} else {
		engine, err = xorm.NewEngine(config.DbType, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=%v&parseTime=true",
			config.DbUser,
			config.DbPassword,
			config.DbHost,
			config.DbPort,
			config.DbName,
			config.DbCharset,
			url.QueryEscape("Asia/Shanghai")))
	}
	engine.ShowSQL(config.DbShowSql)
	return err
}

//New 新建
func New() *xorm.Engine {
	return engine
}
