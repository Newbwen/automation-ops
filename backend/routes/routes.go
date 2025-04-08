package routes

import (
	"github.com/Newbwen/automation-ops/backend/controllers"
	"github.com/Newbwen/automation-ops/backend/utils"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	// 全局中间件：配置 CORS（生产环境应严格限制来源）
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有来源（仅演示用）
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Range")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // 预检请求直接返回204
			return
		}
		c.Next()
	})

	// 认证路由组
	authGroup := r.Group("/")
	{
		// 开放端点
		//authGroup.GET("/api/captcha", controllers.GenerateCaptcha)
		authGroup.POST("/register", controllers.RegisterUser) // 用户注册
		authGroup.POST("/login", controllers.LoginUser)       // 用户登录

		// 需要认证的端点
		authGroup.Use(utils.JWTAuth())                                 // 应用 JWT 中间件
		authGroup.GET("/user", controllers.GetUser)                    // 获取当前用户信息
		authGroup.POST("/change-password", controllers.ChangePassword) // 修改密码
		authGroup.GET("/user/profile", controllers.UserInfo)
	}
	//主机路由组
	hostGroup := r.Group("/host")
	{
		hostGroup.GET("/list", controllers.GetHosts)
		hostGroup.POST("/add", controllers.AddHost)
		hostGroup.PUT("/edit/:id", controllers.EditHost)
		hostGroup.DELETE("/delete/:id", controllers.DeleteHost)
	}

	return r
}
