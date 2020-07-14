package db

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type WhereFun func(db *gorm.DB) *gorm.DB
type FunDB func(db *gorm.DB) error
type FunNum func(db *gorm.DB) error
type FunOrder func(db *gorm.DB) *gorm.DB
type DBHelper struct {

}

type Add struct {
	Err	 error
	FunDB FunDB
}

func (a *Add) Add() *Add {
	if a.FunDB==nil {
		a.Err = errors.New("fundb can not  nil")
		return a
	}
	a.FunDB(MysqlConn)

	
	
	
}
type Update struct {

}
type Delete struct {
	
}
type Status struct {
	
}
type List struct {

}







