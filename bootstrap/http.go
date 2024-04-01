package bootstrap

import (
	"yxkitchen-backend/conf"
	"yxkitchen-backend/pkg/logger"

	"github.com/gin-gonic/gin"
)

func RunWeb() {
	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	// gin 实例
	router := gin.New()

	// 初始化路由绑定
	SetupRoute(router)

	// 运行服务器
	logger.InfoString("bootstrap", "run web server", "http server is running on "+conf.C.Http.Addr)
	if err := router.Run(conf.C.Http.Addr); err != nil {
		logger.FatalString("bootstrap", "run web server", err.Error())
	}
}
