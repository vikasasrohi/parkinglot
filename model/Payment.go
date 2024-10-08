package model

type Payment struct {
	ID     uint
	Amount uint
	Token  *Token
	Method PaymentMethod
}

type PaymentMethod string

const (
	CashPaymentMethod PaymentMethod = "cash"
	CardPaymentMethod PaymentMethod = "card"
	UPIPaymentMethod  PaymentMethod = "upi"
)
