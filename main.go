package main

import (
	"brnnai/producer-sqs/message"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/google/uuid"
)

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func main() {
	log.Println("Starting the producer service")
	awsConfig := GetAWSConfig()
	queueName := aws.String("rawData-SQS")
	message.SQSQueueSetup(awsConfig, queueName)

	user := User{FistName: "Brenon", LastName: "Araujo"}
	var balances []Balance = make([]Balance, 0)
	balance := Balance{
		Group:           "brrn group",
		GroupBalance:    randFloat(999999.90, 999999999.98),
		VariableBalance: randFloat(1000.90, 10000.98),
		ActualBalance:   randFloat(1000.90, 10000.98),
	}
	for i := 0; i < 2; i++ {
		balances = append(balances, balance)
	}
	updateMsg, marshErr := json.Marshal(UpdateMsg{ID: uuid.UUID{}, User: user, Balances: balances})
	if marshErr != nil {
		log.Fatal(marshErr)
	}
	var wg sync.WaitGroup

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			message.SendMessage(message.SQSMessage{Body: updateMsg})
		}()
	}
	wg.Wait()
	fmt.Println("finish")
}
