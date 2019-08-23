package client

import (
	"fmt"
	"time"
)

type ProviderClient interface {
	ValidateTransaction(transaction string) Receipt
}

// Itunes

type ItunesClient struct {
}

func (c ItunesClient) ValidateTransaction(transaction string) Receipt {
	fmt.Printf("Validating Itunes transaction %v ...\n", transaction)
	time.Sleep(time.Second * 1)
	return NewItunesReceipt(transaction, "OK")
}

// Google

type GoogleClient struct {
}

func (c GoogleClient) ValidateTransaction(transaction string) Receipt {
	fmt.Printf("Validating Google transaction %v ...\n", transaction)
	time.Sleep(time.Second * 3)
	return NewGoogleReceipt(transaction, "OK")
}

// Stripe

type StripeClient struct {
}

func (c StripeClient) ValidateTransaction(transaction string) Receipt {
	fmt.Printf("Validating Stripe transaction %v ...\n", transaction)
	time.Sleep(time.Second * 2)
	return NewStripeReceipt(transaction, "OK")
}
