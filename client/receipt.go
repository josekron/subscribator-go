package client

import (
	"encoding/json"
)

// Receipt : receipt which extends Itunes, Google and Stripe receipt
type Receipt interface {
	GetReceipt() string
	GetTransaction() string
	GetProvider() string
}

/////////////
// Itunes  //
/////////////

// ItunesReceipt : contains receipt and latest_receipt_info from Itunes response
type ItunesReceipt struct {
	Provider      string
	Transaction   string
	Receipt       VerifyReceipt `json:"receipt"`
	LatestReceipt []ReceiptInfo `json:"latest_receipt_info"`
}

type VerifyReceipt struct {
	ReceiptType            string        `json:"receipt_type"`
	BundleID               string        `json:"bundle_id"`
	OriginalPurchaseDateMs string        `json:"original_purchase_date_ms"`
	InAppReceipt           []ReceiptInfo `json:"in_app"`
}

type ReceiptInfo struct {
	ProductID              string `json:"product_id"`
	TransactionID          string `json:"transaction_id"`
	OriginalTransactionID  string `json:"original_transaction_id"`
	PurchaseDateMs         string `json:"purchase_date_ms"`
	OriginalPurchaseDateMs string `json:"original_purchase_date_ms"`
	IsTrialPeriod          string `json:"is_trial_period"`
}

func NewItunesReceipt(transaction, receipt string) *ItunesReceipt {

	itunesReceipt := &ItunesReceipt{}
	json.Unmarshal([]byte(receipt), itunesReceipt)

	itunesReceipt.Provider = "itunes"
	itunesReceipt.Transaction = transaction

	return itunesReceipt
}

func (t ItunesReceipt) GetReceipt() string {
	sReceipt, _ := json.Marshal(t.Receipt)
	sLatestReceiptInfo, _ := json.Marshal(t.LatestReceipt)
	return "receipt:" + string(sReceipt) + "\nlatest_receipt_info:" + string(sLatestReceiptInfo)
}

func (t ItunesReceipt) GetTransaction() string {
	return t.Transaction
}

func (t ItunesReceipt) GetProvider() string {
	return t.Provider
}

/////////////
// Google  //
/////////////

type GoogleReceipt struct {
	provider    string
	transaction string
	receipt     string
}

func NewGoogleReceipt(transaction, receipt string) *GoogleReceipt {
	var googleReceipt = &GoogleReceipt{
		provider:    "google",
		transaction: transaction,
		receipt:     receipt,
	}
	return googleReceipt
}

func (t GoogleReceipt) GetReceipt() string {
	return t.receipt
}

func (t GoogleReceipt) GetTransaction() string {
	return t.transaction
}

func (t GoogleReceipt) GetProvider() string {
	return t.provider
}

/////////////
// Stripe  //
/////////////

type StripeReceipt struct {
	provider    string
	transaction string
	receipt     string
}

func NewStripeReceipt(transaction, receipt string) *StripeReceipt {
	var stripeReceipt = &StripeReceipt{
		provider:    "stripe",
		transaction: transaction,
		receipt:     receipt,
	}
	return stripeReceipt
}

func (t StripeReceipt) GetReceipt() string {
	return t.receipt
}

func (t StripeReceipt) GetTransaction() string {
	return t.transaction
}

func (t StripeReceipt) GetProvider() string {
	return t.provider
}
