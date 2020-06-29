package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zngue/go_tool/src/fun/time"
	"github.com/zngue/go_tool/src/sign_chan"
	time2 "time"
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
	fmt.Print()
	db.DB().SetMaxIdleConns(mysql.MaxIdleConns)
	db.DB().SetMaxOpenConns(mysql.MaxOpenConns)
	if mysql.TimeStamp {
		db.Callback().Create().Replace("gorm:update_time_stamp",updateTimeStampForCreateCallback)
		db.Callback().Update().Replace("gorm:update_time_stamp",updateTimeStampForUpdateCallback)
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


type TimeStampModel struct {
	CreatedAt int64 `gorm:"column:created_at" json:"created_at" `
	UpdatedAt int64 `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time2.Time `gorm:"column:deleted_at" json:"deleted_at" sql:"index"`
}
type GormModel struct {
	CreatedAt time2.Time
	UpdatedAt time2.Time
	DeletedAt *time2.Time `sql:"index"`
}