package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var SendPacketDuration = promauto.NewSummary(prometheus.SummaryOpts{
	Name: "sqs_producer_api_send_packet_duration",
	Help: "Send packet messages request duration",
})

var SendRequestTotal = promauto.NewCounter(prometheus.CounterOpts{
	Name: "sqs_producer_api_send_request_total",
	Help: "The total number of requests to  send messages route",
})

var SendRequestDuration = promauto.NewSummary(prometheus.SummaryOpts{
	Name: "sqs_producer_api_send_request_duration",
	Help: "Send messages request duration",
})

var SendBatchRequestTotal = promauto.NewCounter(prometheus.CounterOpts{
	Name: "sqs_producer_api_send_batch_request_total",
	Help: "The total number of send batch requests",
})

var SendBatchRequestDuration = promauto.NewSummary(prometheus.SummaryOpts{
	Name: "sqs_producer_api_send_batch_request_duration",
	Help: "Send messages in batch request duration",
})

var GenMessagesRequestDuration = promauto.NewSummary(prometheus.SummaryOpts{
	Name: "sqs_producer_api_gen_messages_request_duration",
	Help: "Generate messages request duration",
})
