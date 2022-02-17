package api

import (
	"brnnai/producer-sqs/message"
	"brnnai/producer-sqs/utils"
	"brnnai/producer-sqs/worker"
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
			message.SendMessage(message.SQSMessage{Body: utils.GetRandomUpdateMsg()})
		}()
	}
}

func SendSplitsParallel(qtd int) {
	defer worker.CreateWorkerChannels()
	timer := prometheus.NewTimer(GenMessagesRequestDuration)
	genMessages := make([]message.SQSBatchMessage, 0)

}

func SendBatchParallel(qtd int) {
	defer worker.CreateWorkerChannels()
	timer := prometheus.NewTimer(GenMessagesRequestDuration)
	messages := make([]message.SQSBatchMessage, 0)
	for i := 0; i < qtd; i++ {
		id, _ := uuid.NewRandom()
		messages = append(messages, message.SQSBatchMessage{ID: id, Body: utils.GetRandomUpdateMsg()})
	}
	timer.ObserveDuration()
	splits := splitMessages(messages, 10)
	go allocate(splits)
	done := make(chan bool)
	go worker.JobResult(done)
	worker.BatchMessageWorkerPool(500)
	<-done
	log.Printf("Send batch messages in parallel was completed!")
}

func messageGerator() [] {
	for i := 0; i < qtd; i++ {
		id, _ := uuid.NewRandom()
		messages = append(messages, message.SQSBatchMessage{ID: id, Body: utils.GetRandomUpdateMsg()})
	}
}

func splitMessages(msgsToSend []message.SQSBatchMessage, size int) [][]message.SQSBatchMessage {
	splits := make([][]message.SQSBatchMessage, 0)
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

func allocate(splits [][]message.SQSBatchMessage) {
	for _, split := range splits {
		id, _ := uuid.NewRandom()
		job := worker.Job{Id: id, Msgs: split}
		worker.Jobs <- job
	}
	close(worker.Jobs)
}
