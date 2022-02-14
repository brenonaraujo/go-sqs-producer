# go-sqs-producer

A simple app to produce messages to SQS by using aws sdk v2 package for Golang in performatic meaning.

## How to run
 
 | go 1.17 or above is required.

 Just run the command  `go run .` at the root directory

 ## Sending data to SQS

 The service were implemented by using gin as http Router to the api.

 - GET /send/:qtd  Generates and send messages async in parallel
 - GET /send/batch/:qtd  Generates and send messages in batch of 10 messages each, sapaws a hardcoded number of workers into a pool to fan-out messages to sqs.

 ## Metrics

 The server metrics will be at the port 8080, to scrap metrics can simply open the url:

 `http://localhost:8080/metrics`


 ## Performance considerations

 | As I told early, /send/batch route implementes a performatic way to send milions of messages to SQS by using goroutines, channels and an worker pool strategy to send messages in batch and parallel, with this we can produce 1000000 (One milion) messages in less than 180 seconds, thats means we can produce 5500 messages per second with that simple implementation of go routines and channels.

 working in progress...

