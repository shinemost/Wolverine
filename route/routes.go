package route

import (
	"hjfu/Wolverine/controllers"
	"hjfu/Wolverine/setting"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(setting.GinLogger(), setting.GinRecovery(true))
	//注册
	r.POST("/register", controllers.RegisterHandler)
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	return r
}
