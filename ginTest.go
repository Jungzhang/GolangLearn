package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"fmt"
	"net/http"
	"os"
	"io"
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

	router.LoadHTMLGlob("./index.html")
	router.GET("/", urlRoot)

	// GET方法
	// URL参数
	router.GET("/someGet", urlGet)
	// API参数
	router.GET("/someGet/:name/:sex", apiGet)

	// POST方法
	router.POST("/somePost", somePost)

	// 文件上传
	router.POST("/uploadSingleFile", uploadSingFile)

	router.Run(":8080")

}

func urlRoot(t *gin.Context) {
	t.HTML(http.StatusOK, "index.html", gin.H{})
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

func somePost(t *gin.Context) {

	name := t.PostForm("name")
	sex := t.PostForm("sex")

	fmt.Printf("post param name : %v   sex : %v\n", name, sex)
}

func uploadSingFile(c *gin.Context) {

	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		fmt.Println("err :", err)
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	filename := header.Filename
	fmt.Println("filename : " + filename)

	out, err := os.Create(filename)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Println("err :", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.String(http.StatusOK, "upload file successful")
}
