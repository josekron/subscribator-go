package client

import (
	"fmt"
	receipt "subscribator-go/client/receipt"
	"time"
)

type StripeClient struct {
}

func (c StripeClient) ValidateTransaction(transaction string) receipt.IReceipt {
	fmt.Printf("Validating Stripe transaction %v ...\n", transaction)
	time.Sleep(time.Second * 2)
	return receipt.NewStripeReceipt(transaction, "OK")
}
