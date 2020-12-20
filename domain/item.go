package domain

type Item struct {
	Id string
	Title string
	SellerId int32
	BuyingMode string
	Permalink string
	AcceptsMercadopago bool
	Shipping ItemShipping
}

type ItemShipping struct {
	Mode string
}
