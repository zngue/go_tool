package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zngue/go_tool/src/sign_chan"
)
func MysqlConnet()  {
	mysql :=Config.Mysql
	dns:=mysql.Username+":"+mysql.Password+"@tcp("+mysql.Host+":"+mysql.Port+")/"+mysql.DBName+"?charset=utf8&parseTime=True&loc=Asia%2FShanghai"
	db, errDb := gorm.Open("mysql", dns)
	if errDb !=nil {
		sign_chan.SignLog(errDb)
		return
	}
	if mysql.LogMode {
		db = db.LogMode(mysql.LogMode)
	}
	db.DB().SetMaxIdleConns(mysql.MaxIdleConns)
	db.DB().SetMaxOpenConns(mysql.MaxOpenConns)
	MysqlConn = db
	return
}