package types

import "github.com/google/uuid"

type User struct {
	FistName string
	LastName string
	Email    string
}

type Balance struct {
	ActualBalance   float64
	ActualBalance1  float64
	ActualBalance2  float64
	ActualBalance3  float64
	ActualBalance4  float64
	ActualBalance5  float64
	VariableBalance float64
	Group           string
	GroupBalance    float64
	DificultyValue  float64
	User            User
}

type UpdateMsg struct {
	ID       uuid.UUID
	User     User
	Balances []Balance
}
