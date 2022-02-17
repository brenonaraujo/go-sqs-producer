package worker

import (
	"brnnai/producer-sqs/message"
	"sync"
)

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
