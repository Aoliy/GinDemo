package main

import (
	"gindemo/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//Default返回一个默认的路由引擎，里面包含两个中间键logger recovery
	router := gin.Default()
	//.GET() 函数定义了一个简单的路由处理函数
	//router.GET("/ping", func(c *gin.Context) {
	//	//输出json结果给给调用方
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//导入全部的模板到路由引擎
	router.LoadHTMLGlob("templates/*")
	//website分组
	v1 := router.Group("/register") //类似分流
	{
		v1.GET("/", controller.RegisterControllerGet)
		v1.POST("/", controller.RegisterControllerPost)
	}
	v2 := router.Group("/login") //类似分流
	{
		v2.GET("/", controller.LoginControllerGet)
		v2.POST("/", controller.LoginControllerPost)
	}
	v3 := router.Group("/index") //类似分流
	{
		v3.GET("/", controller.IndexControllerGet)
	}

	router.Run()
	//加载静态文件
	//loadStaticFiles()
	//加载动态文件
	//loadFiles()
}

// 路由加载静态文件
func loadStaticFiles() {
	router := gin.Default()
	router.StaticFS("/static", http.Dir("C:\\Users\\92199\\GolandProjects\\gindemo\\public"))
	router.StaticFile("/login", ".\\public\\login.jpg")
	router.Run()
}

// 路由加载静态文件
func loadFiles() {
	router := gin.Default()
	//导入全部的模板到路由引擎
	router.LoadHTMLGlob("templates/*")
	//website分组
	v := router.Group("/123") //类似分流
	{
		v.GET("/index", controller.IndexControllerGet) //url:localhost:8080/123/index
	}
	router.Run()
}
