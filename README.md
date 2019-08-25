# subscribator-go
 
A simple project to handle the subscription renewals from different suppliers (Itunes, Google, Stripe).

### Description ###

It expects a list of Transactions (only transaction ids for now, but each supplier may require additional parameters), validate each transaction id against their supplier and return their receipt.
