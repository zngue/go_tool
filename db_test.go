package main

import (
	"fmt"
	"github.com/zngue/go_tool/src/db"
	"testing"
	"time"
)

func TestBd(t *testing.T)  {
	db.InitDB()
	fmt.Print(db.MysqlConn)
	db.RedisConn.Set("user","zhasnagans",time.Second*10000)
	fmt.Print(db.RedisConn.Get("user").String())
	defer db.ConnClose()
}
