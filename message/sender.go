package message

import (
	"brnnai/producer-sqs/message/sqsApi"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func SendMessage(msg SQSMessage) {
	queue, queueErr := getSqsQueue()
	if queueErr != nil {
		log.Fatal(queueErr)
	}
	bodyJson, bodyErr := msg.Body.MarshalJSON()
	if bodyErr != nil {
		log.Fatal(bodyErr)
	}
	input := &sqs.SendMessageInput{
		DelaySeconds: 10,
		MessageBody:  aws.String(string(bodyJson)),
		QueueUrl:     &queue.QueueURL,
	}
	res, err := sqsApi.SendMsg(context.Background(), queue.client, input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("MessageId: %v, sended!", res.MessageId)
}

func SendBatchMessage(msgs []SQSBatchMessage) {

}

func getSqsQueue() (SQSQueue, error) {
	if Queue.client == nil {
		return Queue, errors.New("sqs Client not setup")
	} else if Queue.QueueURL == "" {
		return Queue, errors.New("queue URL cant be empty or null")
	}
	return Queue, nil
}
