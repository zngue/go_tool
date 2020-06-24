package sign_chan

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)
var SignChan chan os.Signal

func init()  {
	SignChan = make(chan os.Signal)
}
func SignChalNotify()  {
	signal.Notify(SignChan, os.Interrupt)
	<-SignChan
}
func SignLog(err ...interface{})  {
	log.Println(err)
	SignChan<-os.Interrupt
}

type CloseHttp func(ctx context.Context) error
func ListClose( fns ...CloseHttp )  {
	SignChalNotify()
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*5)
	defer cancel()
	if len(fns)>0 {
		for index,fn:=range fns{
			if err :=fn(ctx);err!=nil{
				log.Fatalln(fmt.Sprintf("第%d个强制关闭",index+1))
			}else{
				log.Println(fmt.Sprintf("第%d个关闭成功",index+1))
			}
		}
	}
	log.Println("服务器优雅退出")
}



