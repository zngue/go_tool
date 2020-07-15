package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/db"
)

type FunWhere func(db *gorm.DB) *gorm.DB
type FunDB func(db *gorm.DB) error
type FunNum func(db *gorm.DB) error
type FunOrder func(db *gorm.DB) *gorm.DB
type MapWhere map[string]interface{}
type DBHelper struct {
	Add
	Update
	Delete
	Status
	List
}

type Add struct {
	Err   error
	FunDB FunDB
}

func (a *Add) Add() *Add {
	dbConn := db.MysqlConn
	if a.FunDB == nil {
		a.Err = errors.New("fundb can not  nil")
		return a
	}
	a.Err = a.FunDB(dbConn)
	return a
}

type Update struct {
	Err      error
	FunDB    FunDB
	Data     interface{}
	FunWhere FunWhere
	MapWhere MapWhere
}

func (u *Update) Update() *Update {
	dbConn := db.MysqlConn
	if u.FunWhere != nil {
		dbConn = u.FunWhere(dbConn)
	}
	if u.MapWhere != nil {

	}
	return u
}

type Delete struct {
}
type Status struct {
}
type List struct {
}
