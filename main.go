package main

import (
	"root/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	app := gin.Default()

	route := app
	route.GET("/", func(ctx *gin.Context) {
		isValidated := false
		if !isValidated {
			ctx.AbortWithStatusJSON(4000, gin.H{
				"message": "bad request",
			})
			return
		}
		ctx.JSON(200, gin.H{
			"hello": "world",
		})
	})

	app.Run()

}
