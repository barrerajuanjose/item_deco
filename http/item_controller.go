package http

import (
	"strconv"
	"sync"

	"github.com/barrerajuanjose/item_deco/domain"
	"github.com/barrerajuanjose/item_deco/marshaller"
	"github.com/barrerajuanjose/item_deco/service"
	"github.com/gin-gonic/gin"
)

type ItemController interface {
	Get(ctx *gin.Context)
}

type itemController struct {
	itemMarshaller        marshaller.ItemMarshaller
	itemService           service.ItemService
	userService           service.UserService
	paymentOptionsService service.PaymentOptionsService
}

func NewItemController(itemMarshaller marshaller.ItemMarshaller, itemService service.ItemService, userService service.UserService, paymentOptionsService service.PaymentOptionsService) ItemController {
	return &itemController{
		itemMarshaller:        itemMarshaller,
		itemService:           itemService,
		userService:           userService,
		paymentOptionsService: paymentOptionsService,
	}
}

func (c itemController) Get(ctx *gin.Context) {
	itemId := ctx.Param("item_id")
	buyerId := ctx.Query("buyer_id")

	itemChan := make(chan *domain.Item, 1)
	sellerChan := make(chan *domain.User, 1)
	buyerChan := make(chan *domain.User, 1)
	recommendedPaymentMethodChan := make(chan *domain.PaymentMethod, 1)
	viewChan := make(chan *marshaller.ItemDto, 1)

	var wg sync.WaitGroup
	wg.Add(4)

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

		go func() {
			defer wg.Done()
			defer close(recommendedPaymentMethodChan)
			recommendedPaymentMethodChan <- c.paymentOptionsService.GetRecommendedPaymentMethod(item.Id)
		}()
	}()

	go func() {
		defer close(viewChan)
		wg.Wait()

		itemDomain := <-itemChan
		sellerDomain := <-sellerChan
		buyerDomain := <-buyerChan
		recommendedPaymentMethod := <-recommendedPaymentMethodChan

		viewChan <- c.itemMarshaller.GetView(itemDomain, sellerDomain, buyerDomain, recommendedPaymentMethod)
	}()

	ctx.JSON(200, <-viewChan)
}
