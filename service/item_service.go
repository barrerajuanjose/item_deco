package service

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/barrerajuanjose/item_deco/domain"
	"io/ioutil"
	"net/http"
)

type ItemService interface {
	GetItemById(itemId string) *domain.Item
}

type itemService struct {

}

type itemResponse struct {
	Id string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	SellerId int32 `json:"seller_id,omitempty"`
	BuyingMode string `json:"buying_mode,omitempty"`
	Permalink string `json:"permalink,omitempty"`
	AcceptsMercadopago bool `json:"accepts_mercadopago,omitempty"`
	Shipping ItemResponseShipping
}

type ItemResponseShipping struct {
	Mode string `json:"mode,omitempty"`
}

func NewItemService() ItemService {
	return &itemService{}
}

func (*itemService) GetItemById(itemId string) *domain.Item {
	response, err := http.Get(fmt.Sprintf("https://api.mercadolibre.com/items/%s", itemId))
	if err != nil {
		return nil
	}

	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil
	}

	var itemResponse itemResponse
	_ = json.Unmarshal(respBody, &itemResponse)

	return &domain.Item{
		Id: itemResponse.Id,
		Title: itemResponse.Title,
		SellerId: itemResponse.SellerId,
		BuyingMode: itemResponse.BuyingMode,
		Permalink: itemResponse.Permalink,
		AcceptsMercadopago: itemResponse.AcceptsMercadopago,
		Shipping: domain.ItemShipping{
			Mode: itemResponse.Shipping.Mode,
		},
	}
}
