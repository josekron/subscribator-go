package client

import (
	"fmt"
	"sync"
)

const (
	ITUNES = "itunes"
	GOOGLE = "google"
	STRIPE = "stripe"
)

type ClientService struct {
	itunesClient ItunesClient
	googleClient GoogleClient
	stripeClient StripeClient
}

func NewService(itunesClient ItunesClient, googleClient GoogleClient, stripeClient StripeClient) *ClientService {
	var service = &ClientService{
		itunesClient: itunesClient,
		googleClient: googleClient,
		stripeClient: stripeClient,
	}
	return service
}

/*
ValidateTransactions: receives a list of transactions and validates each transaction with their provider client.

TODO:: Now each Transaction (apple, google, stripe) contains a String transaction but it can be different depending on each supplier.
The same for the clients, they will request to the API of each supplier.
*/
func (cs ClientService) ValidateTransactions(transactions []Transaction) []Receipt {

	receipts := make([]Receipt, len(transactions))
	var wg sync.WaitGroup

	for i, t := range transactions {

		var choosenClient ProviderClient

		switch t.GetProvider() {
		case ITUNES:
			choosenClient = cs.itunesClient
		case GOOGLE:
			choosenClient = cs.googleClient
		case STRIPE:
			choosenClient = cs.stripeClient
		}

		if choosenClient != nil {
			wg.Add(1)
			validateTransaction(choosenClient, t.GetTransaction(), &wg, receipts, i)
		}

	}

	wg.Wait()

	fmt.Printf("Finish to validate %d transactions: \n", len(transactions))

	return receipts
}

func validateTransaction(client ProviderClient, transaction string, wg *sync.WaitGroup, receipts []Receipt, numReceipt int) {

	go func() {
		defer wg.Done()
		receipts[numReceipt] = client.ValidateTransaction(transaction)
	}()
}
