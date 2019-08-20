package client

type StripeReceipt struct {
	provider    string
	transaction string
	receipt     string
}

func NewStripeReceipt(transaction, receipt string) *StripeReceipt {
	var stripeReceipt = &StripeReceipt{
		provider:    "stripe",
		transaction: transaction,
		receipt:     receipt,
	}
	return stripeReceipt
}

func (t StripeReceipt) GetReceipt() string {
	return t.receipt
}

func (t StripeReceipt) GetTransaction() string {
	return t.transaction
}

func (t StripeReceipt) GetProvider() string {
	return t.provider
}
