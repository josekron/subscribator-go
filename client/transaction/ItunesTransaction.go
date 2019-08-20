package client

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
