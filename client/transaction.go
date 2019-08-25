package client

type Transaction interface {
	GetTransaction() string
	GetProvider() string
}

// Itunes

type ItunesTransaction struct {
	provider string
	token    string
}

func NewItunesTransaction(token string) *ItunesTransaction {
	var transaction = &ItunesTransaction{
		provider: "itunes",
		token:    token,
	}
	return transaction
}

func (t ItunesTransaction) GetTransaction() string {
	return t.token
}

func (t ItunesTransaction) GetProvider() string {
	return t.provider
}

// Google

type GoogleTransaction struct {
	provider string
	token    string
}

func NewGoogleTransaction(token string) *GoogleTransaction {
	var transaction = &GoogleTransaction{
		provider: "google",
		token:    token,
	}
	return transaction
}

func (t GoogleTransaction) GetTransaction() string {
	return t.token
}

func (t GoogleTransaction) GetProvider() string {
	return t.provider
}

// Stripe

type StripeTransaction struct {
	provider string
	token    string
}

func NewStripeTransaction(token string) *StripeTransaction {
	var transaction = &StripeTransaction{
		provider: "stripe",
		token:    token,
	}
	return transaction
}

func (t StripeTransaction) GetTransaction() string {
	return t.token
}

func (t StripeTransaction) GetProvider() string {
	return t.provider
}
