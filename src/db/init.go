package db
import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/config"
	"github.com/zngue/go_tool/src/sign_chan"
)
var (
	Config *config.Config
	MysqlConn *gorm.DB
	RedisConn *redis.Client
)
func init()  {
	if Config==nil {
		config.MicroConfig()
		if config.MicroConf!=nil {
			if config.MicroConf.IsMicroConfig {
				Config=config.MicroHttpRequest()
			}else{
				Config=config.YamlToStruck()
			}
		}
	}
	if Config==nil {
		sign_chan.SignLog("配置文件加载失败...")
	}
}

type  AutoDB func(db *gorm.DB)
func InitDB(mysqlDbd ...AutoDB)  {
	//Config =config.JsonToStruck() //获取配置信息
	if Config==nil {
		sign_chan.SignLog("配置文件加载失败...")
		return
	}
	load :=Config.DefaultLoad
	if load.Redis {
		RedisConnet()//链接redis
	}
	if load.Mysql {
		MysqlConnet(mysqlDbd...)//链接mysql数据库
	}
}
//关闭连接池
func ConnClose ()  {
	if Config!=nil {
		if Config.DefaultLoad.Mysql && MysqlConn!=nil {
			MysqlConn.Close()
		}
		if Config.DefaultLoad.Redis && RedisConn!=nil {
			RedisConn.Close()
		}
	}
}

