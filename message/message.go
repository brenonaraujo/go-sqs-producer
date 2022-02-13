package message

import (
	"encoding/json"

	"github.com/google/uuid"
)

type SQSMessage struct {
	Body json.RawMessage
}

type SQSBatchMessage struct {
	id   uuid.UUID
	Body []string
}
