package main

import (
	"gin-study/setting"

	"github.com/gin-gonic/gin"
)

func main() {
	setting.Init()

	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000")
}
