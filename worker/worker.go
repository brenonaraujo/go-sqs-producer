package worker

import (
	"brnnai/producer-sqs/message"
	"log"
	"sync"
)

var Jobs = make(chan Job, 100)
var Results = make(chan Result, 100)

func BatchMessageWorkerPool(qtd int) {
	var wg sync.WaitGroup
	for i := 0; i < qtd; i++ {
		wg.Add(1)
		go batchMessageWorker(&wg)
	}
	wg.Wait()
	close(Results)
}

func batchMessageWorker(wg *sync.WaitGroup) {
	for job := range Jobs {
		if len(job.Msgs) <= 0 {
			continue
		}
		message.SendBatchMessage(job.Msgs)
		output := Result{job, true}
		Results <- output
	}
	wg.Done()
}

func JobResult(done chan bool) {
	for result := range Results {
		log.Printf("Job id %v, finished", result.Job.Id)
	}
	done <- true
}
