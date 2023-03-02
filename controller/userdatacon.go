package controller

import (
	"gindemo/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

// 此结构定义了一个用户，其中包含用户名和密码
type User struct {
	Username string
	Password string
}

// 首页处理器
func IndexControllerGet(c *gin.Context) {
	//username数据拿来
	username := c.Query("username")

	c.HTML(http.StatusOK, "index.html", gin.H{
		"username": username,
	})
}

// 登录处理器
func LoginControllerGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginControllerPost(c *gin.Context) {
	// 从POST请求中读取表单数据
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 创建当前用户
	var user User
	user = User{Username: username, Password: password}

	// 从数据库中查找用户
	db := data.CreateDataBase()
	userdata := data.SelectData1(db, user.Username)
	userdata.Scan(&password)

	// 验证密码是否正确
	if user.Password != password {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// 用户已成功登录
	c.Redirect(http.StatusMovedPermanently, "/index?username="+url.QueryEscape(user.Username))
}

// 注册处理器
func RegisterControllerGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}
func RegisterControllerPost(c *gin.Context) {
	// 从POST请求中读取表单数据
	username := c.PostForm("username")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm-password")

	// 确认密码是否匹配
	if password != confirmPassword {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// 创建新用户
	user := User{Username: username, Password: password}

	// 将新用户保存到数据库
	db := data.CreateDataBase()
	data.CreateUsersTable(db)
	data.InsertData(db, user.Username, user.Password)
	data.CloseDataBase(db)

	// 跳转到登录页面
	c.Redirect(http.StatusMovedPermanently, "/login")
}
