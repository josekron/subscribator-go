# subscribator-go
 
A simple project to handle the subscription renewals from different providers (Itunes, Google, Stripe).

Also, to understand better the use of the goroutines.

### Description ###

It expects a list of transactions (for now, only transaction ids but each provider may require additional parameters), validate each transaction id against their provider and return the response.

### Run ###

ITUNES_URL=https://sandbox.itunes.apple.com/verifyReceipt ITUNES_PWD=XXX go run main.go

### Remaining tasks ###

- Parse String response from Itunes.
- Implement Google client.
- Implement Stripe client.
- Implement endpoints or an interface to comunicate with other services to receive the list of receipts for validation and also return the responses. Options:
  - Gin + protobuffers.
