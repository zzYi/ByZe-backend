package routers

import (
	v1 "ByZe/routers/apis/v1"

	"github.com/gin-gonic/gin"
)

// InitRouter initialization router
func InitRouter() *gin.Engine {
	gin.DisableConsoleColor()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors())
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/ws", v1.WS())
	}

	return r
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 处理请求
		c.Next()
	}
}
