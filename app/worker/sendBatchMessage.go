package worker

import (
	"brnnai/producer-sqs/app/message"
	"sync"
)

func SendBatchMessageWorkerPool(qtd int) {
	var wg sync.WaitGroup
	for i := 0; i < qtd; i++ {
		wg.Add(1)
		go sendBatchMessageWorker(&wg)
	}
	wg.Wait()
	close(Results)
}

func sendBatchMessageWorker(wg *sync.WaitGroup) {
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
