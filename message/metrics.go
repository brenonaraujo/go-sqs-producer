package message

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var MsgSendedTotal = promauto.NewCounter(prometheus.CounterOpts{
	Name: "sqs_producer_message_sended_total",
	Help: "The total number of messages sended to the queue",
})

var SendMessageDuration = promauto.NewSummary(prometheus.SummaryOpts{
	Name: "sqs_producer_send_message_request_duration",
	Help: "Send messages request duration",
})

var MsgSendedPerSec = promauto.NewCounter(prometheus.CounterOpts{
	Name: "sqs_producer_messages_sedend_sec",
	Help: "Messages sended per second",
})
