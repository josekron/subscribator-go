package test

import (
	"fmt"
	client "subscribator-go/client"
	"testing"
)

func TestAppleClient(t *testing.T) {

	itunesClient := client.ItunesClient{}
	itunesTransaction := client.NewItunesTransaction("token_itunes")

	var receipt = itunesClient.ValidateTransaction(itunesTransaction.GetTransaction())

	fmt.Println(receipt)
	if len(receipt.GetReceipt()) == 0 {
		t.Errorf("Expected receipt not empty")
	}
}
