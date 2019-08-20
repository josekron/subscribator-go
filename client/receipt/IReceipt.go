package client

type IReceipt interface {
	GetReceipt() string
	GetTransaction() string
	GetProvider() string
}
