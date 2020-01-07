# subscribator-go
 
A simple project to handle the subscription renewals from different providers (Itunes, Google, Stripe).

Also: 

- Understand better the use of the goroutines.
- Convert string json to struct.

### Description ###

It expects a list of transactions (for now, only transaction ids but each provider may require additional parameters), validate each transaction id against their provider and return the response.

### Run ###

ITUNES_URL=https://sandbox.itunes.apple.com/verifyReceipt ITUNES_PWD=XXX go run main.go

Docker:

docker build -t subscribator-go-docker .

docker run -it -p 8080:8080 subscribator-go-docker

### Remaining tasks ###

- Implement Google client.
- Implement Stripe client.
- Implement endpoints or an interface to comunicate with other services to receive the list of receipts for validation and also return the responses. Options:
  - Gin + protobuffers.
