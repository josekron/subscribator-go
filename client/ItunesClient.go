package client

import (
	"fmt"
	receipt "subscribator-go/client/receipt"
	"time"
)

type ItunesClient struct {
}

func (c ItunesClient) ValidateTransaction(transaction string) receipt.IReceipt {
	fmt.Printf("Validating Itunes transaction %v ...\n", transaction)
	time.Sleep(time.Second * 1)
	return receipt.NewItunesReceipt(transaction, "OK")
}
