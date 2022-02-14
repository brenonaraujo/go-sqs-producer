package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ServerSetup() {
	router := gin.Default()
	router.GET("/send/:quantity", sendByQuantity)
	router.GET("/send/batch/:quantity", sendBatchByQuantity)
	router.GET("/metrics", prometheusHandler())
	router.Run("localhost:8080")
}

func sendByQuantity(c *gin.Context) {
	timer := prometheus.NewTimer(SendRequestDuration)
	defer timer.ObserveDuration()
	qtd := c.Param("quantity")
	qtdValue, err := strconv.Atoi(qtd)
	if err != nil {
		log.Fatal(err)
	}
	SendParallel(qtdValue)
	msg := fmt.Sprintf("Sending %v parallel messages to SQS!", qtd)
	c.IndentedJSON(http.StatusOK, gin.H{"Message": msg})
	SendRequestTotal.Inc()
}

func sendBatchByQuantity(c *gin.Context) {
	timer := prometheus.NewTimer(SendBatchRequestDuration)
	defer timer.ObserveDuration()
	qtd := c.Param("quantity")
	qtdValue, err := strconv.Atoi(qtd)
	if err != nil {
		log.Fatal(err)
	}
	SendBatchParallel(c, qtdValue)
	msg := fmt.Sprintf("Sending %v parallel batch messages to SQS!", qtd)
	c.IndentedJSON(http.StatusOK, gin.H{"Message": msg})
	SendRequestTotal.Inc()
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
