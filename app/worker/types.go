package worker

import (
	"brnnai/producer-sqs/app/message"
	"sync"

	"github.com/google/uuid"
)

type Job struct {
	Id   uuid.UUID
	Msgs []message.SQSMessage
	Msg  message.SQSMessage
}

type Result struct {
	Job        Job
	MsgsResult bool
}

type Worker func(*sync.WaitGroup)
