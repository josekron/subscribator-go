package client

type ITransaction interface {
	GetTransaction() string
	GetProvider() string
}
