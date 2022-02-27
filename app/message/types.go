package message

import (
	"encoding/json"

	"github.com/google/uuid"
)

type SQSMessage struct {
	ID   uuid.UUID
	Body json.RawMessage
}
