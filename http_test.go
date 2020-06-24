package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_tool/src/fun/time"
	"github.com/zngue/go_tool/src/gin_run"
	"testing"
)



func TestHttp2(t *testing.T) {

	s,_:=gin_run.HttpRouterServe("3036", func(group *gin.RouterGroup) {
		
	})
	time.Time()
	s.ListenAndServe()
}









func TestHttp(t *testing.T) {

}
