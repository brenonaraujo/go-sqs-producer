package worker

import (
	"flag"
	"log"
)

var channelCapacity int = *flag.Int("chMaxCap", 100, "Max nunber of parallalel messages at channels job and result")

var Jobs = make(chan Job, channelCapacity)
var Results = make(chan Result, channelCapacity)

func CreateWorkerChannels() {
	Jobs = make(chan Job, channelCapacity)
	Results = make(chan Result, channelCapacity)
}

func JobResult(done chan bool) {
	for result := range Results {
		log.Printf("Job id %v, finished", result.Job.Id)
	}
	done <- true
}
