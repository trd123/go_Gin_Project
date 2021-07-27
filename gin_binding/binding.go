package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	// 验证参数
	Age      int       `form:"age" binding:"required,gt=10"`
	Name     string    `form:"name" binding:"required"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	r := gin.Default()
	r.GET("/bind", func(c *gin.Context) {

		var person Person
		person.Age = 15
		person.Name = "TRD"
		person.Birthday = time.Now()
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, fmt.Sprint(err))
			return
		}
		c.String(200, fmt.Sprintf("%#v", person))
	})
	r.Run()
}
