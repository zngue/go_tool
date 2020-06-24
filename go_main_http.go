package main

/*
import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_tool/app/router"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/gin_run"
	"github.com/zngue/go_tool/src/michttp"
	"github.com/zngue/go_tool/src/sign_chan"
	"sync"
)

type Score struct {
	Times int
}
type Messages struct {
	Data interface{}
	StatusCode int
	Message string

}

func Http(wg *sync.WaitGroup,score *Score,mutex *sync.Mutex)  {
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
	data := map[string]string{
		"userId": "5951",
		"type":   "24",
		"credit": "1",
		"title":  "teste",
	}
	newrequset := michttp.MicHttpRequest{
		Method:    "post",
		EndPoint:  "/v1/user/integral/consume",
		ServiceId: "sy:go:member",
		Param:     data,
		Timeout:   500,
	}
	str ,_:=newrequset.GetRequesBytes()
	var message Messages
	json.Unmarshal(str,&message)
	if message.StatusCode==200 {
		times := score.Times
		times+=1
		score.Times=times
		fmt.Println("-------------",score.Times,"---------")
	}


}
func main()  {
	httpTwo,_:=gin_run.HttpRouterServe("3030", func(group *gin.RouterGroup) {
		router.RouterUserInfo(group)
	})
	httpOne,_:=gin_run.HttpRouterServe("3031", func(group *gin.RouterGroup) {
		router.RouterUserInfo(group)
	})
	httpThree,_:=gin_run.HttpRouterServe("3032", func(group *gin.RouterGroup) {
		router.RouterArticle(group)
	})
	go gin_run.HttpRun(
		func() error {
			return httpOne.ListenAndServe()
		},
		func() error {
			return httpTwo.ListenAndServe()
		},
		func() error {
			return httpThree.ListenAndServe()
		},
	)
	go db.InitDB()
	defer db.ConnClose()
	sign_chan.ListClose(func(ctx context.Context) error {
		return httpOne.Shutdown(ctx)
	}, func(ctx context.Context) error {
		return httpTwo.Shutdown(ctx)
	})
}*/



