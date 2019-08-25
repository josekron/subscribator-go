package client

type Receipt interface {
	GetReceipt() string
	GetTransaction() string
	GetProvider() string
}

// Itunes

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

// Google

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

// Stripe

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
