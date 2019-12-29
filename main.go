package main

//@Author: josekron
//Main class

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	client "subscribator-go/client"
)

func main() {
	fmt.Println("Main - init")

	var itunesURL = os.Getenv("ITUNES_URL")
	var itunesPassword = os.Getenv("ITUNES_PWD")

	itunesClient := client.NewItunesClient(itunesURL, itunesPassword)
	googleClient := client.NewGoogleClient()
	stripeClient := client.NewStripeClient()

	clientService := client.NewService(*itunesClient, *googleClient, *stripeClient)

	//TODO:: Remove it and get transactions from another service
	transactions := []client.Transaction{}
	for i := 0; i < 10; i++ {
		transactionType := rand.Intn(3)

		switch transactionType {
		case 0:
			transactions = append(transactions, client.NewItunesTransaction("token_itunes_"+strconv.FormatInt(int64(i), 10)))
		case 1:
			transactions = append(transactions, client.NewGoogleTransaction("token_google_"+strconv.FormatInt(int64(i), 10)))
		case 2:
			transactions = append(transactions, client.NewStripeTransaction("token_stripe_"+strconv.FormatInt(int64(i), 10)))
		}
	}

	receipts := clientService.ValidateTransactions(transactions)

	for _, receipt := range receipts {
		fmt.Println(receipt.GetTransaction() + " " + receipt.GetReceipt())
	}
}
