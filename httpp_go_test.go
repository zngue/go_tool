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
	"github.com/zngue/go_tool/src/log"
	"github.com/zngue/go_tool/src/michttp"
	"github.com/zngue/go_tool/src/sign_chan"
	"golang.org/x/sync/errgroup"
	log2 "log"
	"sync"
	"time"
)
var (
	g errgroup.Group
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
	//var wg sync.WaitGroup
	////oss.AliOssUploadFile("./110.png","file/1sdasdad2563313.jpeg")
	////times:=time.Time()
	////timess := time.TimeToFormat(times,"2006.01.02 15:04:05")
	//gin.SetMode(gin.ReleaseMode)
	//
	//
	//log.LogInfo("ok")
	//wg.Wait()
	////defer db.ConnClose()
	//log.LogInfo("ok")
	ser1,_:=gin_run.HttpRouterServe("3030",func(group *gin.RouterGroup) {
		router.RouterUserInfo(group)
	})
	serv2,_:=gin_run.HttpRouterServe("3031", func(group *gin.RouterGroup) {
		router.IdeaRouter(group)
	})
	db.InitDB()
	err1 := ser1.ListenAndServe()
	err2:=serv2.ListenAndServe()
	fmt.Println(err1,err2)
	defer db.ConnClose()
	sign_chan.SignChalNotify()
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*5)
	defer cancel()
	g.Go(func() error {
		return ser1.Shutdown(ctx)
	})
	g.Go(func() error {
		return serv2.Shutdown(ctx)
	})
	if err :=g.Wait();err!=nil{
		log2.Fatalln("服务器关闭")
	}
	log.LogInfo("服务器优雅退出1")


}


*/