package main

import (
	"brnnai/producer-sqs/api"
	"brnnai/producer-sqs/message"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func main() {
	log.Println("Starting the SQS producer API")
	awsConfig := GetAWSConfig()
	queueName := aws.String("rawData-SQS")
	message.SQSQueueSetup(awsConfig, queueName)
	api.SetupAPI()
	log.Println("wating for messages to be sended!")
}
