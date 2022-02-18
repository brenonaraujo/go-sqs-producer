package worker

import (
	"brnnai/producer-sqs/message"
	"sync"
)

func SendMessageWorkerPool(qtd int) {
	var wg sync.WaitGroup
	for i := 0; i < qtd; i++ {
		wg.Add(1)
		go sendMessageWorker(&wg)
	}
	wg.Wait()
	close(Results)
}

func sendMessageWorker(wg *sync.WaitGroup) {
	for job := range Jobs {
		message.SendMessage(job.Msg)
		output := Result{job, true}
		Results <- output
	}
	wg.Done()
}
