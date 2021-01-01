package service

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/barrerajuanjose/item_deco/domain"
)

type PaymentOptionsService interface {
	GetRecommendedPaymentMethod(item_id string) *domain.PaymentMethod
}

type paymentOptionsService struct {
}

type paymentOptionsResponse struct {
	PaymentMethods []paymentMethodResponse `json:"payment_methods,omitempty"`
}

type paymentMethodResponse struct {
	Id string `json:"id,omitempty"`
}

func NewPaymentOptionsService() PaymentOptionsService {
	return &paymentOptionsService{}
}

func (*paymentOptionsService) GetRecommendedPaymentMethod(itemId string) *domain.PaymentMethod {
	response, err := http.Get(fmt.Sprintf("https://api.mercadolibre.com/items/%s/payment_options?status=active", itemId))
	if err != nil {
		return nil
	}

	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil
	}

	var paymentOptionsResponse paymentOptionsResponse
	_ = json.Unmarshal(respBody, &paymentOptionsResponse)

	return &domain.PaymentMethod{
		Id: paymentOptionsResponse.PaymentMethods[0].Id,
	}
}
