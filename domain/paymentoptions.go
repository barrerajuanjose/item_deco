package domain

type PaymentOptions struct {
	PaymentMethods []PaymentMethod
}

type PaymentMethod struct {
	Id string
}
