package api

import (
	"brnnai/producer-sqs/message"
	"brnnai/producer-sqs/types"
	"encoding/json"
	"log"
	"math/rand"
	"sync"

	"github.com/google/uuid"
)

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func genRandomUpdateMsg() []byte {
	user := types.User{FistName: "Brenon", LastName: "Araujo"}
	var balances []types.Balance = make([]types.Balance, 0)
	balance := types.Balance{
		Group:           "brrn group",
		GroupBalance:    randFloat(999999.90, 999999999.98),
		VariableBalance: randFloat(1000.90, 10000.98),
		ActualBalance:   randFloat(1000.90, 10000.98),
	}
	for i := 0; i < 2; i++ {
		balances = append(balances, balance)
	}
	updateMsg, marshErr := json.Marshal(types.UpdateMsg{ID: uuid.UUID{}, User: user, Balances: balances})
	if marshErr != nil {
		log.Fatal(marshErr)
	}
	return updateMsg
}

func SendParallel(qtd int) {
	var wg sync.WaitGroup
	for i := 0; i < qtd; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			message.SendMessage(message.SQSMessage{Body: genRandomUpdateMsg()})
		}()
	}
}

func SendBatchParallel(qtd int) {

}
