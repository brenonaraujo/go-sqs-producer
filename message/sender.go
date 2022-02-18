package message

import (
	"brnnai/producer-sqs/message/sqsApi"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/prometheus/client_golang/prometheus"
)

func SendMessage(msg SQSMessage) {
	timer := prometheus.NewTimer(SendMessageDuration)
	queue, queueErr := GetSqsQueue()
	if queueErr != nil {
		log.Fatal(queueErr)
	}
	input := &sqs.SendMessageInput{
		DelaySeconds: 10,
		MessageBody:  aws.String(string(msg.Body)),
		QueueUrl:     &queue.QueueURL,
	}
	_, err := sqsApi.SendMsg(context.Background(), queue.client, input)
	if err != nil {
		log.Fatal(err)
	}
	MsgSendedTotal.Inc()
	timer.ObserveDuration()
}

func SendBatchMessage(msgs []SQSMessage) {
	timer := prometheus.NewTimer(SendBatchMessageDuration)
	defer timer.ObserveDuration()
	queue, queueErr := GetSqsQueue()
	if queueErr != nil {
		log.Fatal(queueErr)
	}
	entries := make([]types.SendMessageBatchRequestEntry, 0)
	for _, msg := range msgs {
		entries = append(entries, types.SendMessageBatchRequestEntry{
			Id:          aws.String(msg.ID.String()),
			MessageBody: aws.String((string(msg.Body))),
		})
	}
	input := &sqs.SendMessageBatchInput{
		QueueUrl: &queue.QueueName,
		Entries:  entries,
	}
	_, err := sqsApi.SendBatchMsg(context.Background(), queue.client, input)
	if err != nil {
		log.Fatal(err)
	}
	MsgBatchSendedTotal.Inc()
}
