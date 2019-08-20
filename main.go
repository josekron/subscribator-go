package main

//@Author: josekron
//Main class

import (
	"fmt"
	"math/rand"
	"strconv"
	client "subscribator-go/client"
	transaction "subscribator-go/client/transaction"
)

func main() {
	fmt.Println("Main - init")

	clientService := client.NewService(client.ItunesClient{}, client.GoogleClient{}, client.StripeClient{})

	transactions := []transaction.ITransaction{}
	for i := 0; i < 10; i++ {
		transactionType := rand.Intn(3)

		switch transactionType {
		case 0:
			transactions = append(transactions, transaction.NewItunesTransaction("token_itunes_"+strconv.FormatInt(int64(i), 10)))
		case 1:
			transactions = append(transactions, transaction.NewGoogleTransaction("token_google_"+strconv.FormatInt(int64(i), 10)))
		case 2:
			transactions = append(transactions, transaction.NewStripeTransaction("token_stripe_"+strconv.FormatInt(int64(i), 10)))
		}
	}

	receipts := clientService.ValidateTransactions(transactions)

	for _, receipt := range receipts {
		fmt.Println(receipt.GetTransaction() + " " + receipt.GetReceipt())
	}
}
