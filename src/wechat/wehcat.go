package wechat

import (
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/sign_chan"
)
var WeChat *wechat.Wechat
func init()  {
	if WeChat!=nil {
		ChatInit()
	}
}
func ChatInit()  {
	chatConfig :=db.Config.WeChat
	if &chatConfig==nil {
		sign_chan.SignLog("we chat config err")
	}
	redisConfig:=db.Config.Redis
	if &redisConfig==nil {
		sign_chan.SignLog("redis config err")
	}
	redisCache :=cache.NewRedis(&cache.RedisOpts{
		Host: redisConfig.Host,
		Password: redisConfig.Password,
		Database: redisConfig.DBNum,
	})
	WeChat=wechat.NewWechat(&wechat.Config{
		AppID:          chatConfig.AppID,
		AppSecret:     chatConfig.AppSecret,
		Token:          chatConfig.EncodingAESKey,
		EncodingAESKey: chatConfig.EncodingAESKey,
		PayMchID:		chatConfig.PayMchID,
		PayNotifyURL:	chatConfig.PayNotifyURL,
		PayKey:			chatConfig.PayKey,
		Cache:          redisCache,
	})
}
