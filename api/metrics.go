package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var SendRequestTotal = promauto.NewCounter(prometheus.CounterOpts{
	Name: "sqs_producer_api_send_total",
	Help: "The total number of requests to  send messages route",
})

var SendRequestDuration = promauto.NewSummary(prometheus.SummaryOpts{
	Name: "sqs_producer_send_request_duration",
	Help: "Send messages request duration",
})
