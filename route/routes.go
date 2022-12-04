package route

import (
	"github.com/gin-gonic/gin"
	"hjfu/Wolverine/controllers"
	"hjfu/Wolverine/logger"
	"hjfu/Wolverine/middleware"
	"net/http"
)

// 路由
func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(mode)
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")

	//注册
	v1.POST("/register", controllers.RegisterHandler)
	v1.POST("/login", controllers.LoginHandler)
	v1.GET("/community", controllers.CommunityHandler)

	v1.GET("/ping", middleware.JWTAuthMiddleWare(), controllers.PingHandler)

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})
	return r
}
