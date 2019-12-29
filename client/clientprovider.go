package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// ProviderClient : interface implemented by each client
type ProviderClient interface {
	ValidateTransaction(transaction string) Receipt
}

/////////////
// Itunes  //
/////////////

// ItunesClient : client Itunes
type ItunesClient struct {
	url        string
	password   string
	httpClient http.Client
}

// NewItunesClient : function to get a new pointer to itunes client
func NewItunesClient(url, password string) *ItunesClient {
	if len(url) == 0 {
		url = "https://sandbox.itunes.apple.com/verifyReceipt"
	}
	var itunesClient = &ItunesClient{
		url:        url,
		password:   password,
		httpClient: http.Client{},
	}
	return itunesClient
}

// ValidateTransaction : validate an Itunes transaction receipt
func (c ItunesClient) ValidateTransaction(transaction string) Receipt {
	fmt.Printf("Validating Itunes transaction %v...%v\n", transaction[0:5], transaction[len(transaction)-5:])

	var request = ItunesRequest{transaction, c.password}
	json, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return NewItunesReceipt(transaction, "500")
	}

	req, err := http.NewRequest(http.MethodPost, c.url, bytes.NewBuffer(json))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return NewItunesReceipt(transaction, "500")
	}

	fmt.Println(resp.StatusCode)

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
		return NewItunesReceipt(transaction, bodyString)
	}

	return NewItunesReceipt(transaction, strconv.Itoa(resp.StatusCode))
}

type ItunesRequest struct {
	ReceiptData string `json:"receipt-data"`
	Password    string `json:"password"`
}

/////////////
// Google  //
/////////////

// GoogleClient : client Google
type GoogleClient struct {
}

// NewGoogleClient : function to get a new pointer to Google client
func NewGoogleClient() *GoogleClient {
	var googleClient = &GoogleClient{}
	return googleClient
}

// ValidateTransaction : validate an Google transaction receipt
func (c GoogleClient) ValidateTransaction(transaction string) Receipt {
	fmt.Printf("Validating Google transaction %v ...\n", transaction)
	time.Sleep(time.Second * 3)
	return NewGoogleReceipt(transaction, "OK")
}

/////////////
// Stripe  //
/////////////

// StripeClient : client Stripe
type StripeClient struct {
}

// NewStripeClient : function to get a new pointer to Stripe client
func NewStripeClient() *StripeClient {
	var stripeClient = &StripeClient{}
	return stripeClient
}

// ValidateTransaction : validate an Stripe transaction receipt
func (c StripeClient) ValidateTransaction(transaction string) Receipt {
	fmt.Printf("Validating Stripe transaction %v ...\n", transaction)
	time.Sleep(time.Second * 2)
	return NewStripeReceipt(transaction, "OK")
}
