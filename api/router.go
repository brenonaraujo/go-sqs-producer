package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupAPI() {
	router := gin.Default()
	router.GET("/send/:quantity", sendByQuantity)
	router.GET("/metrics", prometheusHandler())
	router.Run("localhost:8080")
}

func sendByQuantity(c *gin.Context) {
	qtd := c.Param("quantity")
	qtdValue, err := strconv.Atoi(qtd)
	if err != nil {
		log.Fatal(err)
	}
	wg := SendParallel(qtdValue)
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "Parallel message send started"})
	wg.Done()
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
