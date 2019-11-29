package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
)

func Init() {
	//数据库配置
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")

	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8mb4"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	//注册数据库
	orm.RegisterDataBase("default", "mysql", dsn, 10, 600)
	// 自动建表
	orm.RunSyncdb("default", false, true)
	//sql命令
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}
