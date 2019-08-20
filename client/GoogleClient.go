package client

import (
	"fmt"
	receipt "subscribator-go/client/receipt"
	"time"
)

type GoogleClient struct {
}

func (c GoogleClient) ValidateTransaction(transaction string) receipt.IReceipt {
	fmt.Printf("Validating Google transaction %v ...\n", transaction)
	time.Sleep(time.Second * 3)
	return receipt.NewGoogleReceipt(transaction, "OK")
}
