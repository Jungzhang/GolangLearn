package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"fmt"
)

func main() {

	/*
		gin的默认用法,一个空服务
	*/
	/*
	router := gin.Default()
	router.Run()
	*/

	/*
		创建出带有默认中间件的路由：日志和恢复中间件
	*/
	router := gin.Default()

	/*
		创建不带中间件的路由
	*/
	//router := gin.New()

	/*
		自定义HTTP的配置
	*/
	/*s := &http.Server{
		Addr:              "8080",
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	s.ListenAndServe()
	*/

	// URL参数
	router.GET("someGet", urlGet)
	// API参数
	router.GET("someGet/:name/:sex", apiGet)
	router.Run(":8080")

}

func urlGet(t *gin.Context) {

	name := t.Query("name")
	sex := t.Query("sex")

	fmt.Printf("url param name : %v   sex : %v\n", name, sex)
}

func apiGet(t *gin.Context) {
	// API参数
	name := t.Param("name")
	sex := t.Param("sex")

	fmt.Printf("api param name : %v   sex : %v\n", name, sex)
}
