package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"log"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {

		t := time.Now()

		// content中设置值
		c.Set("example", 1234)

		// 发送request之前
		c.Next()

		// 发送request之后
		latency := time.Since(t)
		log.Println(latency)

		// 通过ResponseWrite获取状态
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {

	router := gin.Default()

	router.Use(MiddleWare())
	router.GET("/logger", func(context *gin.Context) {
		example := context.MustGet("example")
		log.Println(example)
		example, err := context.Get("example")
		if err == false {
			log.Println("the key example is not exist")
			return
		}
		log.Println(example)
	})

	router.Run(":8080")

}
