package message

import (
	"brnnai/producer-sqs/message/sqsApi"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func SendMessage(msg SQSMessage) {
	queue, queueErr := GetSqsQueue()
	if queueErr != nil {
		log.Fatal(queueErr)
	}
	log.Printf("Start to send message to queue: %v.\n", queue.QueueName)
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
	log.Printf("Message: %v sended!\n", *res.MessageId)
}

func SendBatchMessage(msgs []SQSBatchMessage) {

}
