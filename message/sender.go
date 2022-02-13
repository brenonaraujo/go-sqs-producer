package message

import (
	"brnnai/producer-sqs/message/sqsApi"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/prometheus/client_golang/prometheus"
)

func SendMessage(msg SQSMessage) {
	timer := prometheus.NewTimer(SendMessageDuration)
	defer timer.ObserveDuration()
	queue, queueErr := GetSqsQueue()
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
	_, err := sqsApi.SendMsg(context.Background(), queue.client, input)
	if err != nil {
		log.Fatal(err)
	}
	MsgSendedTotal.Inc()
}

func SendBatchMessage(msgs []SQSBatchMessage) {

}
