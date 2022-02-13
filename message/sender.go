package message

import (
	"brnnai/producer-sqs/message/sqsApi"
	"brnnai/producer-sqs/metrics"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/prometheus/client_golang/prometheus"
)

func SendMessage(msg SQSMessage) {
	timer := prometheus.NewTimer(metrics.SendMessageDuration)
	queue, queueErr := GetSqsQueue()
	if queueErr != nil {
		log.Fatal(queueErr)
	}
	//log.Printf("Start to send message to queue: %v.\n", queue.QueueName)
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
	defer timer.ObserveDuration()
	metrics.MsgSendedPerSec.Inc()
	// log.Printf("Message: %v sended!\n", *res.MessageId)
}

func SendBatchMessage(msgs []SQSBatchMessage) {

}
