package main

import (
	"brnnai/producer-sqs/message"
	"encoding/json"
	"log"
	"math/rand"

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

	user := User{FistName: "Brenon", LastName: "Araujo", Email: "brenonaraujo@gmail.com"}
	var balances []Balance = make([]Balance, 0)
	balance := Balance{
		Group:           "brrn group",
		GroupBalance:    randFloat(999999.90, 999999999.98),
		VariableBalance: randFloat(1000.90, 10000.98),
		ActualBalance:   randFloat(1000.90, 10000.98),
		DificultyValue:  randFloat(0.10, 10.00),
	}
	for i := 0; i < 10; i++ {
		balances = append(balances, balance)
	}
	updateMsg, marshErr := json.Marshal(UpdateMsg{ID: uuid.UUID{}, User: user, Balances: balances})
	if marshErr != nil {
		log.Fatal(marshErr)
	}
	message.SendMessage(message.SQSMessage{Body: updateMsg})
}
