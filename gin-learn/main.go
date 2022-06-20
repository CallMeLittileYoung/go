package main

import (
	"gin-learn/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//模拟db db阿33
var db = make(map[string]string)

func person(context *gin.Context) {
	var person entity.Person
	if context.ShouldBind(&person) == nil {
		log.Println(person.Address)
		log.Println(person.Birthday)
		log.Println(person.Birthday)
		context.JSON(http.StatusOK, "success")
	} else {
		context.JSON(http.StatusOK, "fail")
	}
}
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
	//用 AsciiJSON 方法可以生成只包含 ASCII 字符的 JSON 格式数据，对于非 ASCII 字符会进行转义：
	r.GET("/asciiJson", func(context *gin.Context) {
		data := map[string]interface{}{
			"name": "叫我小年轻",
			"bbb":  "br",
		}
		context.AsciiJSON(http.StatusOK, data)
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
	//公共这种方法来注册一个路由挺好的
	r.POST("/person", person)
	r.GET("/person", person)
	return r
}
func main() {
	router := setRouter()
	err := router.Run(":10086")
	if err != nil {
		return
	}
}
