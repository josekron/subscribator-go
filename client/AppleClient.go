package client

import (
	"fmt"
	receipt "subscribator-go/client/receipt"
	"time"
)

type AppleClient struct {
}

func (c AppleClient) ValidateTransaction(transaction string) receipt.IReceipt {
	fmt.Printf("Validating  Apple transaction %v ...\n", transaction)
	time.Sleep(time.Second * 1)
	return receipt.NewAppleReceipt(transaction, "OK")
}
