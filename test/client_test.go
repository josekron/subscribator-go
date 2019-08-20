package test

import (
	"fmt"
	client "subscribator-go/client"
	transaction "subscribator-go/client/transaction"
	"testing"
)

func TestAppleClient(t *testing.T) {

	appleClient := client.AppleClient{}
	appleTransaction := transaction.NewAppleTransaction("token_apple")

	var receipt = appleClient.ValidateTransaction(appleTransaction.GetTransaction())

	fmt.Println(receipt)
	if len(receipt.GetReceipt()) == 0 {
		t.Errorf("Expected receipt not empty")
	}
}
