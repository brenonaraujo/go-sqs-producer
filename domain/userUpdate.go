package domain

import "github.com/google/uuid"

type User struct {
	FistName string
	LastName string
	Email    string
}
type Balance struct {
	ActualBalance   float64
	VariableBalance float64
	Group           string
	GroupBalance    float64
	DificultyValue  float64
}
type UpdateMsg struct {
	ID       uuid.UUID
	User     User
	Balances []Balance
}
