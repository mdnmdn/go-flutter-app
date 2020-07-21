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
			"ts":      time.Now().Format("2"),
		})
	})
	r.Run(":8003")
}
