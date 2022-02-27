package main

import (
	"brnnai/producer-sqs/app/api"
	"brnnai/producer-sqs/app/message"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func main() {
	log.Println("Starting the SQS producer API")
	awsConfig := GetAWSConfig()
	message.SQSQueueSetup(awsConfig, aws.String("rawData-SQS"))
	api.ServerSetup()
	log.Println("wating for messages to be sended!")
}
