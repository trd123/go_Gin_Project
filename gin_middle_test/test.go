package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义中间件
func myTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	// 统计时间
	since := time.Since(start)
	fmt.Println("程序用时：", since)
}

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 注册中间件
	r.Use(myTime)
	{
		shopingGroup := r.Group("/shoping")
		{
			shopingGroup.GET("/index", shopIndexHandler)
			shopingGroup.GET("/home", shopHomeHandler)
		}
		r.Run()
	}
}

func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
}

func shopHomeHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
}
