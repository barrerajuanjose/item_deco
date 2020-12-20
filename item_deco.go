package main

import (
	"github.com/barrerajuanjose/item_deco/http"
	"github.com/barrerajuanjose/item_deco/marshaller"
	"github.com/barrerajuanjose/item_deco/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	itemController := http.NewItemController(marshaller.NewItemMarshaller(), service.NewItemService(), service.NewUserService())

	r.GET("/items/:item_id", itemController.Get)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}