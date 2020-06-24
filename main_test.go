package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_tool/app/router"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/gin_run"
	"github.com/zngue/go_tool/src/sign_chan"
	"testing"
)


func TestMainTe(t *testing.T) {

	httpOne,_:=gin_run.HttpRouterServe("3031", func(group *gin.RouterGroup) {
		router.RouterArticle(group)
	})
	httpTwo,_:=gin_run.HttpRouterServe("3032", func(group *gin.RouterGroup) {
		router.IdeaRouter(group)
	})
	go gin_run.HttpRun(func() error {
		return httpOne.ListenAndServe()
	})
	go gin_run.HttpRun(func() error {
		return httpTwo.ListenAndServe()
	})
	go db.InitDB()
	defer db.ConnClose()
	sign_chan.ListClose(
		func(ctx context.Context) error {
			return httpOne.Shutdown(ctx)
		}, func(ctx context.Context) error {
			return  httpTwo.Shutdown(ctx)
		})
}



