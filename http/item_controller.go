package http

import (
	"github.com/barrerajuanjose/item_deco/domain"
	"github.com/barrerajuanjose/item_deco/marshaller"
	"github.com/barrerajuanjose/item_deco/service"
	"github.com/gin-gonic/gin"
	"sync"
)

type ItemController interface {
	Get(ctx *gin.Context)
}

type itemController struct {
	itemMarshaller marshaller.ItemMarshaller
	itemService service.ItemService
	userService service.UserService
}

func NewItemController(itemMarshaller marshaller.ItemMarshaller, itemService service.ItemService, userService service.UserService) ItemController {
	return &itemController{
		itemMarshaller: itemMarshaller,
		itemService: itemService,
		userService: userService,
	}
}

func (c itemController) Get(ctx *gin.Context) {
	itemId := ctx.Param("item_id")
	var itemDomain *domain.Item

	itemChan := make(chan *domain.Item, 1)
	sellerChan := make(chan *domain.User, 1)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(itemChan)
		itemChan <- c.itemService.GetItemById(itemId)
	}()

	go func() {
		defer wg.Done()
		defer close(sellerChan)
		itemDomain = <- itemChan
		sellerChan <- c.userService.GetUserById(itemDomain.SellerId)
	}()

	wg.Wait()

	sellerDomain := <- sellerChan

	ctx.JSON(200, c.itemMarshaller.GetView(itemDomain, sellerDomain))
}