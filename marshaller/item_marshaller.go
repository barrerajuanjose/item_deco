package marshaller

import "github.com/barrerajuanjose/item_deco/domain"

type ItemDto struct {
	Title                    string `json:"title,omitempty"`
	Permalink                string `json:"permalink,omitempty"`
	SellerNickname           string `json:"seller_nickname,omitempty"`
	BuyerNickname            string `json:"buyer_nickname,omitempty"`
	ShippingMode             string `json:"shipping_mode,omitempty"`
	BestPaymentMethod        string `json:"best_payment_method,omitempty"`
	BestShippingMethod       string `json:"best_shipping_method,omitempty"`
	RecommendedPaymentMethod string `json:"recommended_payment_method,omitempty"`
}

type ItemMarshaller interface {
	GetView(item *domain.Item, seller *domain.User, buyer *domain.User, recommendedPaymentMethod *domain.PaymentMethod) *ItemDto
}

type itemMarshaller struct {
}

func NewItemMarshaller() ItemMarshaller {
	return &itemMarshaller{}
}

func (m itemMarshaller) GetView(item *domain.Item, seller *domain.User, buyer *domain.User, recommendedPaymentMethod *domain.PaymentMethod) *ItemDto {
	var buyerNickname string
	if buyer != nil {
		buyerNickname = buyer.Nickname
	}

	return &ItemDto{
		Title:                    item.Title,
		Permalink:                item.Permalink,
		SellerNickname:           seller.Nickname,
		BuyerNickname:            buyerNickname,
		RecommendedPaymentMethod: recommendedPaymentMethod.Id,
	}
}
