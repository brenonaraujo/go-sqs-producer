package message

import (
	"brnnai/producer-sqs/message/sqsApi"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSQueue struct {
	QueueName string
	QueueURL  string
	client    *sqs.Client
}

var Queue SQSQueue

func SQSQueueSetup(cfg aws.Config, queueName *string) {
	client := sqs.NewFromConfig(cfg)
	input := &sqs.GetQueueUrlInput{
		QueueName: queueName,
	}
	result, err := sqsApi.GetQueueURL(context.TODO(), client, input)
	if err != nil {
		log.Fatal(err)
	}
	Queue.client = client
	Queue.QueueName = *queueName
	Queue.QueueURL = *result.QueueUrl
}
