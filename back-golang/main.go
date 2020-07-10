package main

import (
	"database/sql"
	"finance/common"
	_ "finance/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
)

func init(){
	conf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		logs.Error("new conf failed ,err : %v", err)
		return
	}
	var dbSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		conf.String("mysql_user"),conf.String("mysql_password"),conf.String("mysql_host"),
		conf.String("mysql_port"),conf.String("mysql_db"))
	logs.Debug(dbSourceName)
	if db, err := sql.Open("mysql", dbSourceName);err != nil {
		logs.Error("db err :%v",err)
	} else {
		var emailStmpPort,emailLimit,emailExpire int
		if emailStmpPort,err = conf.Int("email_stmp_port");err != nil {
			panic("conf err - email_stmp_port err : "+ err.Error())
		}
		if emailLimit,err = conf.Int("email_limit");err != nil {
			panic("conf err - email_limit err : "+ err.Error())
		}
		if emailExpire,err = conf.Int("email_expire");err != nil {
			panic("conf err - email_expire err : "+ err.Error())
		}
		common.Source = &common.Conf{
			Db:db,
			EmailStmpHost :conf.String("email_stmp_host"),
			EmailStmpPort :emailStmpPort,
			EmailFrom :conf.String("email_from"),
			EmailPassword :conf.String("email_password"),
			EmailToers :conf.String("email_toers"),
			EmailCcers :conf.String("email_ccers"),
			EmailLimit:emailLimit,
			EmailExpire:emailExpire,
		}
		logs.Info("init conf success!")
	}
	common.Init()
}
func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//允许访问所有源
		AllowAllOrigins: true,
		//可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		//其中Options跨域复杂请求预检
		AllowMethods:   []string{"*"},
		//指的是允许的Header的种类
		AllowHeaders: 	[]string{"*"},
		//公开的HTTP标头列表
		ExposeHeaders:	[]string{"Content-Length"},
		//如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))
	beego.SetLevel(beego.LevelWarning)
	beego.Run()
}

