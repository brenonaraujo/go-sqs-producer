# go-sqs-producer

A simple app to produce messages to sqs by using aws sdk v2 package for Golang.

## How to run
 
 | go 1.17 or above is required.

 Just run the command  `go run .` at the root directory

 ## Sending data to SQS

 The service were implemented by using gin as http Router to the api.

 - GET /send/:qtd  Generates and send messages async in parallel
 - GET /send/batch/:qtd  Generates and send messages in batch of 10 messages each, sapaws a hardcoded number of workers into a pool to fan-out messages to sqs.

 ## Metrics

 The server metrics will be at the port 2112, to scrap metrics can simply open the url:

 `http://localhost:8080/metrics`

 working in progress...

