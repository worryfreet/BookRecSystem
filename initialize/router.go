package initialize

import (
	_ "BookRecSystem/docs"
	"BookRecSystem/global"
	"BookRecSystem/middleware"
	"BookRecSystem/router"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.New()
	Router.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	// 为用户头像和文件提供静态地址
	Router.StaticFS("/api/"+global.GSD_CONFIG.Local.Path, http.Dir(global.GSD_CONFIG.Local.Path))
	global.GSD_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.GSD_LOG.Info("use middleware cors")
	Router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GSD_LOG.Info("register swagger handler")

	// 获取路由组实例
	systemRouter := router.RouterGroupApp.System
	bookRecRouter := router.RouterGroupApp.BookRec
	PublicGroup := Router.Group("api")
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		bookRecRouter.InitBookPubRecRouter(PublicGroup)
	}
	PrivateGroup := Router.Group("api")
	PrivateGroup.Use(middleware.JWTAuth())
	//PrivateGroup.Use(middleware.CasbinHandler())
	{
		systemRouter.InitUserRouter(PrivateGroup)
		bookRecRouter.InitBookRecPriRouter(PrivateGroup)
	}
	global.GSD_LOG.Info("Router Register Success")
	return Router
}
