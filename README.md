# subscribator-go
 
A simple project to handle the subscription renewals from different providers (Itunes, Google, Stripe).

### Description ###

It expects a list of transactions (for now, only transaction ids but each provider may require additional parameters), validate each transaction id against their provider and return their receipt.
