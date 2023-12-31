package routes

import (
	"superman-gin/app/controllers/app"
	"superman-gin/app/controllers/common"
	"superman-gin/app/middleware"
	"superman-gin/app/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/auth/register", app.Register)
	router.POST("/auth/login", app.Login)
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", app.Info)
		authRouter.POST("/auth/logout", app.Logout)
		authRouter.POST("/image_upload", common.ImageUpload)
	}
	router.GET("/auth/helloworld", app.Helloworld)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
