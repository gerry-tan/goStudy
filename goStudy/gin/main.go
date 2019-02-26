package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type UserInfo struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func testHandle(context *gin.Context) {
	var userInfo UserInfo
	example := context.MustGet("example").(string)
	log.Print(example)
	if err := context.ShouldBind(&userInfo); err == nil {
		//结构体自动封装为 json 返回
		context.JSON(http.StatusOK, userInfo)
	} else {
		//手动拼接 json
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func StatCost() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()

		//可以设置一些公共参数
		context.Set("example", "12345")
		//等其他中间件先执行
		context.Next()
		//获取耗时
		latency := time.Since(t)
		log.Printf("cost time: %v", latency)
	}
}

func main() {
	engine := gin.Default()
	engine.Use(StatCost())

	engine.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	engine.DELETE("/delete/:id", func(context *gin.Context) {
		id := context.Param("id")
		context.JSON(http.StatusOK, gin.H{"deleteId": id})
	})

	engine.POST("/test", testHandle)

	engine.POST("/loginJson", func(context *gin.Context) {
		var userInfo UserInfo
		if err := context.ShouldBindJSON(&userInfo); err == nil {
			context.JSON(http.StatusOK, gin.H{
				"user":     userInfo.User,
				"password": userInfo.Password,
			})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})

	engine.POST("/loginForm", func(context *gin.Context) {
		var userInfo UserInfo
		if err := context.ShouldBind(&userInfo); err == nil {
			context.JSON(http.StatusOK, gin.H{
				"user":     userInfo.User,
				"password": userInfo.Password,
			})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})

	engine.Run(":8080")
}
