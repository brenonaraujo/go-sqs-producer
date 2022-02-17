package worker

import "flag"

var channelCapacity int = *flag.Int("chMaxCap", 500, "Max nunber of parallalel messages at channels job and result")

var Jobs = make(chan Job, channelCapacity)
var Results = make(chan Result, channelCapacity)

func CreateWorkerChannels() {
	Jobs = make(chan Job, channelCapacity)
	Results = make(chan Result, channelCapacity)
}
