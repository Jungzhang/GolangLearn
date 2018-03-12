package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"log"
	"net/http"
	"fmt"
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

// 鉴权中间件, 鉴定session_id的合法性
func authMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookies, err := c.Request.Cookie("session_id"); err == nil {
			value := cookies.Value
			fmt.Println(value)
			if value == "123" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
}

func main() {

	router := gin.Default()

	// 只针对/before这个路径注册, 其他路径不起作用, 所以将下边的Use注释掉则Use后的语句将报错
	router.GET("/before", MiddleWare(), func(c *gin.Context) {
		example := c.MustGet("example")
		log.Println(example)
		c.JSON(http.StatusOK, gin.H{
			"example": example,
		})
	})

	// 注册全局
	router.Use(MiddleWare())
	router.GET("/middleware", func(context *gin.Context) {
		example := context.MustGet("example")
		log.Println(example)
		example, err := context.Get("example")
		if err == false {
			log.Println("the key example is not exist")
			return
		}
		log.Println(example)
	})

	// 因为上边注册了全局中间件, 所以这里可以直接使用
	router.GET("/testMiddleware", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"example": c.MustGet("example")})
	})

	// 模拟登陆的页面
	router.GET("/login", func(c *gin.Context) {
		cookie := &http.Cookie{
			Name:     "session_id",
			Value:    "123",
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		c.String(http.StatusOK, "Login successful")
	})

	// 注册鉴权中间件, 模拟登陆后对保护页面的鉴权(session_id)
	router.GET("/home", authMiddleWare(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Already enter home page"})
	})

	router.Run(":8080")

}
