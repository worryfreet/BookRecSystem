package system

import (
	v1 "BookRecSystem/api"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (b *UserRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	userRouter := Router.Group("user")
	var userApi = v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.POST("login", userApi.Login)       // 登录
		userRouter.POST("register", userApi.Register) // 注册
	}
	return userRouter
}
func (b *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	userRouter := Router.Group("user")
	var userApi = v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.GET("userList", userApi.UserList)                     // 后台获取用户列表
		userRouter.PUT("updatePassword", userApi.UpdatePassword)         // 更新密码
		userRouter.GET("userInfo", userApi.UserInfo)                     // 获取用户信息
		userRouter.PUT("updateUserInterest", userApi.UpdateUserInterest) // 修改用户信息
	}
	return userRouter
}
