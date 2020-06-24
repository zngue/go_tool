package gin_run

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_tool/src/db"
	"net/http"
)

type RouterFun func(group *gin.RouterGroup)

func Routers(fn RouterFun,ginType int) *gin.Engine {
	gin.SetMode("debug")
	var router *gin.Engine
	if ginType==1 {
		router = gin.Default()
	}else {
		router = gin.New()
	}
	ApiGroup := router.Group("")
	fn(ApiGroup)
	return router
}
type FnHttp func(group *gin.Engine)

func HttpRouterServe(port string,fnRouter RouterFun) (*http.Server,*gin.Engine) {
	gin.SetMode("debug")
	router := gin.New()
	//router.Use(gin.Logger())
	//router.Use(gin.Recovery())
	ApiGroup := router.Group("")
	fnRouter(ApiGroup)
	return Http(router,port),router
}
//test
func InitROuter  (Router *gin.RouterGroup )  {
	test:=Router.Group("test")
	{
		test.GET("index", func(context *gin.Context) {
			//context.String(200,"te")
			context.JSON(200,db.Config)
		})
	}
}
