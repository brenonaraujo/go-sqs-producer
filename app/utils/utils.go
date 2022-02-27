package utils

import (
	"brnnai/producer-sqs/app/types"
	"encoding/json"
	"log"
	"math/rand"

	"github.com/google/uuid"
)

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func GetRandomData(size int) []byte {
	user := types.User{FistName: "Brenon", LastName: "Araujo"}
	var balances []types.Balance = make([]types.Balance, 0)
	for i := 0; i < size; i++ {
		balance := types.Balance{
			Group:           "brnnai group",
			GroupBalance:    randFloat(999999.90, 999999999.98),
			VariableBalance: randFloat(10000.90, 10000000.98),
			ActualBalance:   randFloat(10000.90, 10000000.98),
			ActualBalance1:  randFloat(10000.90, 10000000.98),
			ActualBalance2:  randFloat(10000.90, 10000000.98),
			ActualBalance3:  randFloat(10000.90, 10000000.98),
			ActualBalance5:  randFloat(10000.90, 10000000.98),
		}
		balances = append(balances, balance)
	}
	uuid, _ := uuid.NewRandom()
	updateMsg, marshErr := json.Marshal(types.UpdateMsg{ID: uuid, User: user, Balances: balances})
	if marshErr != nil {
		log.Fatal(marshErr)
	}
	return updateMsg
}
