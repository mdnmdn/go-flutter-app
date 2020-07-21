package backend

import (
	"time"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
			"ts":      time.Now().Format("2006-01-02T15:04:05-0700"),
		})
	})
	r.Run(":8003")
}
