package client

type ItunesReceipt struct {
	provider    string
	transaction string
	receipt     string
}

func NewItunesReceipt(transaction, receipt string) *ItunesReceipt {
	var itunesReceipt = &ItunesReceipt{
		provider:    "itunes",
		transaction: transaction,
		receipt:     receipt,
	}
	return itunesReceipt
}

func (t ItunesReceipt) GetReceipt() string {
	return t.receipt
}

func (t ItunesReceipt) GetTransaction() string {
	return t.transaction
}

func (t ItunesReceipt) GetProvider() string {
	return t.provider
}
