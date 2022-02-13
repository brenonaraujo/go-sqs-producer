package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var MsgSended = promauto.NewCounter(prometheus.CounterOpts{
	Name: "sqs_producer_message_sended_total",
	Help: "The total number of messages sended to the queue",
})
