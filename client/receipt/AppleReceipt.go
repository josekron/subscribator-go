package client

type AppleReceipt struct {
	provider    string
	transaction string
	receipt     string
}

func NewAppleReceipt(transaction, receipt string) *AppleReceipt {
	var appleReceipt = &AppleReceipt{
		provider:    "apple",
		transaction: transaction,
		receipt:     receipt,
	}
	return appleReceipt
}

func (t AppleReceipt) GetReceipt() string {
	return t.receipt
}

func (t AppleReceipt) GetTransaction() string {
	return t.transaction
}

func (t AppleReceipt) GetProvider() string {
	return t.provider
}
