package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_"github.com/degary/usermanage/routers"
	"github.com/astaxie/beego"
	"github.com/degary/usermanage/services"
	_ "github.com/go-sql-driver/mysql"
	"github.com/degary/usermanage/modules"
	_"github.com/astaxie/beego/session/redis"
)

func init() {
	setupLogger()
	if err := setupDb(); err !=nil {
		return
	}
}


func main() {

	beego.Run()
}

func setupDb() error {
	db_name := beego.AppConfig.String("database::DbName")
	db_type := beego.AppConfig.String("database::DbType")
	db_host := beego.AppConfig.String("database::DbHost")
	db_port := beego.AppConfig.String("database::DbPort")
	db_username :=beego.AppConfig.String("database::DbUsername")
	db_password := beego.AppConfig.String("database::DbPassword")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&charset=utf8mb4",
		db_username,
		db_password,
		db_host,
		db_port,
		db_name,
		)

	orm.RegisterDriver(db_type,orm.DRMySQL)
	err := orm.RegisterDataBase("default", db_type, dsn)
	if err != nil {
		logs.Error("connet db err: %v\n",err)
		return err
	}else {
		logs.Info("connect mysql success")
	}

	//注册表
	orm.RegisterModel(&modules.User{}) //指针类型
	orm.RegisterModel(&modules.Article{})
	orm.RunSyncdb("default",false,true) //同步数据库
	//结构体对应表是否存在
	//表不存在,则创建对应的表
	//表存在,属性是否在表中存在
	//属性不存在,添加列
	//索引是否存在,索引不存在,添加索引
	if user := services.GetUserByName("admin");user == nil {
		logs.Info("创建管理员账号")
		services.AddUser("admin", "abc@123","beijing")

	}
	return nil
}

func setupLogger()  {
	logs.SetLogger(logs.AdapterFile,`{"filename":"usermanage.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
}

func setupSession()  {
}

