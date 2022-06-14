package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//模拟db
var db = make(map[string]string)

func setRouter() *gin.Engine {
	//初始化
	r := gin.Default()

	r.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})
	//新增一个路径
	r.GET("/user/:name", func(context *gin.Context) {
		name := context.Params.ByName("name")
		value, ok := db[name]
		if ok {
			context.JSON(http.StatusOK, gin.H{"user": name, "value": value})
		} else {
			context.JSON(http.StatusOK, gin.H{"user": name, "status": "no value"})
		}
	})
	// 需要 HTTP 基本授权认证的子路由群组设置
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":   "bar", // 用户名:foo 密码:bar
		"manu":  "123", // 用户名:manu 密码:123
		"young": "little",
	}))
	// 保存用户信息路由
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// 解析并验证 JSON 格式请求数据
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})
	return r
}
func main() {
	router := setRouter()
	router.Run(":10086")
}
