package client

type AppleTransaction struct {
	provider string
	token    string
}

func NewAppleTransaction(token string) *AppleTransaction {
	var transaction = &AppleTransaction{
		provider: "apple",
		token:    token,
	}
	return transaction
}

func (t AppleTransaction) GetTransaction() string {
	return t.token
}

func (t AppleTransaction) GetProvider() string {
	return t.provider
}
