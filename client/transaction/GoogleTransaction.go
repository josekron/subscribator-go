package client

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
