package tools

import (
	"dingdongser/common"
	models "dingdongser/models"
	"fmt"
	"xorm.io/xorm"
)
import _ "github.com/go-sql-driver/mysql"
type Orm struct {
	*xorm.Engine
}
var DbEngine *Orm
func InitDbEngine(cfg *Config){
	/*
	1.加载配置文件
	2.创建数据库表格
	3.赋值给全局变量
	*/
	dbConfig := cfg.Database
	conn := dbConfig.DbUser + ":" + dbConfig.DbPass + "@/" + dbConfig.DbName + "?charset=" + dbConfig.CharSet
	fmt.Println("连接字符：",conn)
	engine,err :=xorm.NewEngine(dbConfig.Drive,conn)
	common.ErrHandler("连接失败",err)
	_ = engine.Sync2(new(models.User),new(models.Password),new(models.Userproblem),new(models.Organization))
	orm :=new(Orm)
	orm.Engine = engine
	DbEngine = orm
}
