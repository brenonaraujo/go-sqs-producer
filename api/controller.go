package api

import (
	"brnnai/producer-sqs/message"
	"brnnai/producer-sqs/utils"
	"brnnai/producer-sqs/worker"
	"log"
	"sync"

	"github.com/google/uuid"
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

func SendBatchParallel(qtd int) {
	genMessages := make([]message.SQSBatchMessage, 0)
	for i := 0; i < qtd; i++ {
		id, _ := uuid.NewRandom()
		genMessages = append(genMessages, message.SQSBatchMessage{ID: id, Body: utils.GetRandomUpdateMsg()})
	}
	splits := splitMessages(genMessages)
	go allocate(splits)
	done := make(chan bool)
	go worker.JobResult(done)
	worker.BatchMessageWorkerPool(100)
	<-done
	log.Printf("Send batch messages in parallel was completed!")
	worker.Results = make(chan worker.Result, 100)
	worker.Jobs = make(chan worker.Job, 100)
}

func splitMessages(msgsToSend []message.SQSBatchMessage) [][]message.SQSBatchMessage {
	splits := make([][]message.SQSBatchMessage, 0)
	size := 10
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
