package http

import (
	"github.com/barrerajuanjose/item_deco/domain"
	"github.com/barrerajuanjose/item_deco/marshaller"
	"github.com/barrerajuanjose/item_deco/service"
	"github.com/gin-gonic/gin"
	"strconv"
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
	buyerId := ctx.Query("buyer_id")

	itemChan := make(chan *domain.Item, 1)
	sellerChan := make(chan *domain.User, 1)
	buyerChan := make(chan *domain.User, 1)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		defer close(buyerChan)
		if buyerInt, err := strconv.ParseInt(buyerId, 10, 32); err == nil {
			buyerChan <- c.userService.GetUserById(int32(buyerInt))
		} else {
			buyerChan <- nil
		}
	}()

	go func() {
		defer wg.Done()
		defer close(itemChan)
		item := c.itemService.GetItemById(itemId)
		itemChan <- item

		go func() {
			defer wg.Done()
			defer close(sellerChan)
			sellerChan <- c.userService.GetUserById(item.SellerId)
		}()
	}()

	wg.Wait()

	itemDomain := <- itemChan
	sellerDomain := <- sellerChan
	buyerDomain := <- buyerChan

	ctx.JSON(200, c.itemMarshaller.GetView(itemDomain, sellerDomain, buyerDomain))
}