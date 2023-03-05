package api

import (
	"brnnai/producer-sqs/app/message"
	"brnnai/producer-sqs/app/utils"
	"brnnai/producer-sqs/app/worker"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
)

func SendParallel(qtd int) {
	var wg sync.WaitGroup
	for i := 0; i < qtd; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			message.SendMessage(message.SQSMessage{Body: utils.GetRandomData(30)})
		}()
	}
}

// without compression 210 is packets will be < 64kb
// 60 workers was the number of best performance with that package size
func SendPacketsParallel(qtd int) {
	defer worker.CreateWorkerChannels()
	timer := prometheus.NewTimer(GenMessagesRequestDuration)
	timer.ObserveDuration()
	messages := make([]message.SQSMessage, 0)
	for i := 0; i < qtd; i++ {
		messages = append(messages, message.SQSMessage{Body: utils.GetRandomData(30)})
	}
	timer.ObserveDuration()
	go allocateMessageJobs(messages)
	done := make(chan bool)
	go worker.JobResult(done)
	worker.SendMessageWorkerPool(100)
	<-done
	log.Printf("Send parallel packets as messages completed!!")
}

/**
The param messageSize defines the number of itens in the array that`s from now is used to define the message size
Reference table:
	30 = 10kb
	20 = 7kb
	10 = 3kb
	5  = 1.87kb
**/
func SendBatchParallel(qtd int, messageSize int) {
	defer worker.CreateWorkerChannels()
	timer := prometheus.NewTimer(GenMessagesRequestDuration)
	messages := make([]message.SQSMessage, 0)
	for i := 0; i < qtd; i++ {
		id, _ := uuid.NewRandom()
		messages = append(messages, message.SQSMessage{ID: id, Body: utils.GetRandomData(messageSize)})
	}
	timer.ObserveDuration()
	splits := splitMessages(messages, 10)
	go allocateBatchJobs(splits)
	done := make(chan bool)
	go worker.JobResult(done)
	worker.SendBatchMessageWorkerPool(100)
	<-done
	log.Printf("Send batch messages in parallel completed!")
}

func splitMessages(msgsToSend []message.SQSMessage, size int) [][]message.SQSMessage {
	splits := make([][]message.SQSMessage, 0)
	var end int
	for i := 0; i <= len(msgsToSend); i += size {
		end += size
		if end > len(msgsToSend) {
			end = len(msgsToSend)
		}
		splits = append(splits, msgsToSend[i:end])
	}
	return splits
}

func allocateBatchJobs(splits [][]message.SQSMessage) {
	for _, split := range splits {
		id, _ := uuid.NewRandom()
		job := worker.Job{Id: id, Msgs: split}
		worker.Jobs <- job
	}
	close(worker.Jobs)
}

func allocateMessageJobs(messages []message.SQSMessage) {
	for _, message := range messages {
		id, _ := uuid.NewRandom()
		job := worker.Job{Id: id, Msg: message}
		worker.Jobs <- job
	}
	close(worker.Jobs)
}
