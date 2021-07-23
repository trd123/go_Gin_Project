package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var wg sync.WaitGroup

func main() {
	r := gin.Default()
	/*
		1.gin支持加载HTML模板, 然后根据模板参数进行配置并返回相应的数据，
			本质上就是字符串替换
	*/
	// LoadHTMLGlob() 方法加载模板文件
	r.LoadHTMLGlob("./*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "gin.html", gin.H{"title": "我是测试", "ce": "123456"})
	})

	// 2.重定向
	r.GET("/gin", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
	})

	// 3.同步异步
	// 3.1 异步
	r.GET("/unsync", func(c *gin.Context) {
		// 需要一个副本
		copyContext := c.Copy()
		// 异步处理

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				str := strconv.Itoa(i)
				defer wg.Done()
				time.Sleep(1 * time.Second)
				fmt.Printf("异步执行：" + str + "次" + copyContext.Request.URL.Path + "\n")
			}()
			wg.Wait()
		}

	})
	// 3.2 同步
	r.GET("/sync", func(c *gin.Context) {
		for i := 0; i < 10; i++ {
			time.Sleep(3 * time.Second)
			fmt.Printf("同步执行：" + strconv.Itoa(i) + "次" + c.Request.URL.Path + "\n")
		}

	})

	r.Run()
}
