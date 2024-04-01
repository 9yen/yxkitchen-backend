// Package routes 注册路由
package routes

import (
	controllers "yxkitchen-backend/app/http/controllers/api/v1"
	"yxkitchen-backend/app/http/middlewares"
	"yxkitchen-backend/conf"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	var v1 *gin.RouterGroup
	if conf.C.App.ApiDomain == "" {
		v1 = r.Group("/api/v1")
	} else {
		v1 = r.Group("/v1")
	}

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("200-H"))
	{
		authGroup := v1.Group("/auth")
		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 测试时，可以调高一点
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			uc := new(controllers.UsersController)

			// 获取当前用户
			usersGroup := v1.Group("/users")
			{
				usersGroup.GET("", uc.Index)
			}
		}
	}
}
