package message

import (
	"brnnai/producer-sqs/message/sqsApi"
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSQueue struct {
	QueueName string
	QueueURL  string
	client    *sqs.Client
}

var queue SQSQueue

func GetSqsQueue() (SQSQueue, error) {
	if queue.client == nil {
		return queue, errors.New("sqs Client not setup")
	} else if queue.QueueURL == "" {
		return queue, errors.New("queue URL cant be empty or null")
	}
	return queue, nil
}

func SQSQueueSetup(cfg aws.Config, queueName *string) {
	log.Printf("Setup queue %v to be used.\n", *queueName)
	client := sqs.NewFromConfig(cfg)
	input := &sqs.GetQueueUrlInput{
		QueueName: queueName,
	}
	result, err := sqsApi.GetQueueURL(context.TODO(), client, input)
	if err != nil {
		log.Fatal(err)
	}
	queue.client = client
	queue.QueueName = *queueName
	queue.QueueURL = *result.QueueUrl
	log.Printf("Queue: %v succesfully setup!\n", *queueName)
}
