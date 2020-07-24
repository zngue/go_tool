package test

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/helper"
	"testing"
)

type UserInfo struct {

	ID int `gorm:"primary_key;auto_increment;column:id"`
	Name string

}

type Request struct {
	ExtType string `json:"ext_type"`

}




type Response struct {
	Add UserInfo
	List []*UserInfo
	Delete UserInfo
	ListMore []*UserInfo
	Status UserInfo
	Update UserInfo
}
func TestDbHelp(t *testing.T) {
	db.InitDB(func(db *gorm.DB) {
		db.AutoMigrate(new(UserInfo))
	})
	var req Response
	listHeler:=helper.ListMore(&helper.Helper{
		FunDoDB: func(db *gorm.DB) error {
			return db.Find(&req.List).Error
		},
	},&helper.Helper{
		FunDoDB: func(db *gorm.DB) error {
			return db.Find(&req.ListMore).Error
		},
	});
	c:=helper.Model{
		TypeExt: "7",
		IsGroupRun: true,
		ListMore:listHeler,
	}
	c.Run()

	fmt.Println(req)

	fmt.Println(c)





















}
