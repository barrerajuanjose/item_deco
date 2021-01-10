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

	itemController := http.NewItemController(marshaller.NewItemMarshaller(), service.NewItemService(), service.NewUserService(), service.NewPaymentOptionsService())

	r.GET("/items/:item_id", itemController.Get)

	r.Run(":5000")
}
