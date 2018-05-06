package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	一个有趣的问题
	问题：该程序运行时会panic
	问题分析：因为这里出现了二义性，在/user/:id中因为这里需要从url中传入参数，
			所以后边的/user/status救护产生二义性，比如id是一个string类型
			又恰好是status，这时候就不知道到底调用的是那个接口了
*/

func main() {
	router := gin.Default()
	router.POST("/user/:id", fun1)
	router.POST("/user/status", func2)

	router.Run(":8080")
}

func fun1(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func func2(c *gin.Context) {
	c.String(http.StatusOK, "World")
}
