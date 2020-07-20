package helper

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/db"
	"sync"
)

type FunWhere func(db *gorm.DB) *gorm.DB
type FunDoDB func(db *gorm.DB) error
type FunNum func(db *gorm.DB) error
type FunOrder func(db *gorm.DB) *gorm.DB
type FunField func(db *gorm.DB) *gorm.DB
type FunPreload func(db *gorm.DB) *gorm.DB
type MapWhere map[string]interface{}
type DbArr func(db ...FunDoDB)

type Helper struct {
	MapWhere MapWhere
	FunDoDB FunDoDB
	FunWhere FunWhere
	FunOrder FunOrder
	FunNum FunNum
	FunField FunField
	FunPreload FunPreload
	Request *request.Page
	DBConn *gorm.DB
	Err error
	Model interface{}
}
func (h *Helper) GoRun(wg *sync.WaitGroup)  {
	defer wg.Done()
	h.Run()
}

func (h *Helper) Run ()  *Helper {
	dbConn := db.MysqlConn
	if h.Model!=nil {
		dbConn = dbConn.Model(&h.Model)
	}
	if h.MapWhere!=nil {
		dbConn = h.MapWhereDeal(dbConn)
	}
	if h.FunWhere!=nil {
		dbConn = h.FunWhere(dbConn)
	}
	if h.FunOrder!=nil {
		dbConn  = h.FunOrder(dbConn)
	}
	if h.FunField!=nil {
		dbConn = h.FunField(dbConn)
	}
	if h.FunPreload!=nil {
		dbConn = h.FunPreload(dbConn)
	}
	if h.Request!=nil && h.Request.IsCount!=0 && h.FunNum!=nil  {
		h.Err = h.FunNum(dbConn)
		if h.Err!=nil {
			return h
		}
		if h.Request.IsCount==2 {
			return h
		}
	}
	if h.Request!=nil {
		if h.Request.PageSize>0 && h.Request.Page>0 {
			dbConn = dbConn.Offset(h.Request.PageLimitOffset()).Limit(h.Request.PageSize)
		}
	}
	if h.FunDoDB!=nil {
		if err :=h.FunDoDB(dbConn);err!=nil {
			h.Err=err
			return h
		}
	}
	h.DBConn = dbConn
	return h
}
func (h *Helper) MapWhereDeal(dbConn *gorm.DB) *gorm.DB {
	if h.MapWhere!=nil {
		if h.MapWhere != nil {
			for key,val :=range h.MapWhere{
				dbConn = dbConn.Where(key,val)
			}
		}
	}
	return  dbConn
}
