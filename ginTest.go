package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"fmt"
	"net/http"
	"os"
	"io"
	"log"
	"time"
	"github.com/gin-gonic/gin/binding"
)

type user struct {
	Name string `form:"name" json:"name" binding:"required"`
	Sex  string `form:"sex" json:"sex" binding:"required"`
	Age  int    `form:"age" json:"age"`
	Stu  student
}

type student struct {
	Class string `form:"class" json:"class" binding:"required"`
	Grade string `form:"grade" json:"grade"`
}

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

	// 文件上传(单个文件)
	router.POST("/uploadSingleFile", uploadSingFile)

	// 文件上传(多个文件)
	router.POST("/uploadMultiFile", uploadMultiFile)

	// 返回JSON示例
	router.POST("/testJSON", testJSON)

	// 参数绑定(JSON)
	router.POST("/testBindJSON", testBindJSON)

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
		log.Println("err :", err)
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
		log.Println("err :", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.String(http.StatusOK, "upload file successful")
}

func uploadMultiFile(c *gin.Context) {

	err := c.Request.ParseMultipartForm(102400)
	if err != nil {
		log.Println(err)
	}

	formData := c.Request.MultipartForm

	files := formData.File["upload"]
	for i := range files {

		file, err := files[i].Open()
		// 注册defer, 在这个循环中实际上是注册了多次defer
		// 这里的写法不是最优的,因为defer在return的时候才会调用
		// 所以在该func return前会造成fd的浪费, 从而降低并发
		// 可以将这一部分内容封装成func,从而提高并发
		defer file.Close()
		if err != nil {
			log.Println(err)
		}

		out, err := os.Create(files[i].Filename)
		defer out.Close()
		if err != nil {
			log.Println(err)
		}

		_, err = io.Copy(out, file)
		if err != nil {
			log.Println(err)
		}

		time.Sleep(time.Second * 3)

		c.String(http.StatusCreated, "upload successful   ")
	}
}

func testJSON(c *gin.Context) {

	name := c.PostForm("name")
	sex := c.PostForm("sex")
	age := c.PostForm("age")

	class := c.PostForm("class")
	grade := c.PostForm("grade")

	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"sex":  sex,
		"age":  age,
		"student": gin.H{
			"class": class,
			"grade": grade,
		}})
}

func testBindJSON(c *gin.Context) {

	var user user
	var err error

	contentType := c.Request.Header.Get("Content-Type")

	fmt.Println("context/type : " + contentType)

	switch contentType {
	case "application/json":
		err = c.BindJSON(&user)
	case "application/x-www-form-urlencoded":
		err = c.BindWith(&user, binding.Form)
	}

	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"name": user.Name,
		"age":  user.Age,
		"sex":  user.Sex,
		"studen":
		gin.H{
			"class": user.Stu.Class,
			"grade": user.Stu.Grade,
		}})
}
