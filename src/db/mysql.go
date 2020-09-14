package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zngue/go_tool/src/fun/time"
	"github.com/zngue/go_tool/src/sign_chan"
)
func MysqlConnet(mysqlDB ...AutoDB)  {
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
	if len(mysqlDB)>0 {
		for _,fn:=range mysqlDB{
			fn(db)
		}
	}
	db.DB().SetMaxIdleConns(mysql.MaxIdleConns)
	db.DB().SetMaxOpenConns(mysql.MaxOpenConns)
	if mysql.TimeStamp {
		db.Callback().Create().Replace("gorm:update_time_stamp",updateTimeStampForCreateCallback)
		db.Callback().Update().Replace("gorm:update_time_stamp",updateTimeStampForUpdateCallback)
	}
	if mysql.Prefix!="" {
		gorm.DefaultTableNameHandler= func(db *gorm.DB,defaultTableName string) string{
			return mysql.Prefix+defaultTableName
		}
	}
	MysqlConn = db
	return
}
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdatedAt", time.Time())
	}
}
// // 注册新建钩子在持久化之前
func updateTimeStampForCreateCallback(scope *gorm.Scope) {

	if !scope.HasError() {
		nowTime := time.Time()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}


