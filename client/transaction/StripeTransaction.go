package client

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
