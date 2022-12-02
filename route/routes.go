package route

import (
	"hjfu/Wolverine/controllers"
	"hjfu/Wolverine/logger"
	"hjfu/Wolverine/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//注册
	r.POST("/register", controllers.RegisterHandler)
	r.POST("/login", controllers.LoginHandler)
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})
	r.GET("/ping", middleware.JWTAuthMiddleWare(), controllers.PingHandler)

	return r
}
