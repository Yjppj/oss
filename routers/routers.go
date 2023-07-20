package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"sso/controller"
	"sso/initialize/logger"
	"sso/middlewares"
)

func Setup(mode string) *gin.Engine {

	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //设置为发布模式
	}
	//con参数检验 server逻辑处理 dao数据操作
	r := gin.New()
	//r.Use(cors.Default()) //第三方库
	r.Use(middlewares.Cors())
	fmt.Println(middlewares.Cors())

	zap.L().Info("跨域配置完成")

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	Ding := r.Group("/marchsoft")
	{
		Ding.POST("login", controller.LoginHandler)
		Ding.POST("ticket", controller.LoginHandler)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
