package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"hjfu/Wolverine/controllers"
	_ "hjfu/Wolverine/docs" // 千万不要忘了导入把你上一步生成的docs
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
	v1.GET("/community/:id", controllers.CommunityDetailHandler)
	v1.POST("/post", middleware.JWTAuthMiddleWare(), controllers.CreatePostHandler)
	v1.GET("/post/:id", controllers.PostDetailHandler)
	v1.GET("/postList", controllers.GetPostListHandler)
	v1.POST("/like/", middleware.JWTAuthMiddleWare(), controllers.PostLikeHandler)
	// 最新或者最热列表
	v1.GET("/postList2", controllers.GetPostListHandler2)

	v1.GET("/ping", middleware.JWTAuthMiddleWare(), controllers.PingHandler)

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	//swagger 接口文档
	// http://localhost:你的端口号/swagger/index.html 可以看到接口文档
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	return r
}
