package client

import receipt "subscribator-go/client/receipt"

type IClient interface {
	ValidateTransaction(transaction string) receipt.IReceipt
}
