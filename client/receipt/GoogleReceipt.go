package client

type GoogleReceipt struct {
	provider    string
	transaction string
	receipt     string
}

func NewGoogleReceipt(transaction, receipt string) *GoogleReceipt {
	var googleReceipt = &GoogleReceipt{
		provider:    "google",
		transaction: transaction,
		receipt:     receipt,
	}
	return googleReceipt
}

func (t GoogleReceipt) GetReceipt() string {
	return t.receipt
}

func (t GoogleReceipt) GetTransaction() string {
	return t.transaction
}

func (t GoogleReceipt) GetProvider() string {
	return t.provider
}
