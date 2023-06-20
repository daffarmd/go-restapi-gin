package main

import (
	// "fmt"
	"root/controllers/productcontroller"
	"root/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Print("test")
	r := gin.Default()
	models.ConnectToDB()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	r.Run(":4477")
}
