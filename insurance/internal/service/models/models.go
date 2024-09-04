package models

import "time"

const (
	Pending = iota
	Active
	Payed
	Canceled
)

const (
	GroupInsurance = iota
	GroupPayoff
)

const (
	MonthPeriod    = time.Hour * 24 * 30
	SixMonthPeriod = MonthPeriod * 6
	YearPeriod     = MonthPeriod * 12
)

type Insurance struct {
	ID         string    `json:"id" bson:"id"`
	SenderAddr string    `json:"senderAddr" bson:"senderAddr"`
	Amount     int64     `json:"amount" bson:"amount"`
	CarImage   []byte    `json:"carImage" bson:"-"`
	ActiveTill time.Time `json:"activeTill" bson:"activeTill"`
}

type Payoff struct {
	CarImage     []byte  `json:"carImage"`
	ReceiverAddr string  `json:"receiverAddr"`
	Multiplier   float32 `json:"multiplier"`
}

type InsuranceData struct {
	ID                 string    `json:"id"`
	Status             int       `json:"status"`
	ActiveTill         time.Time `json:"activeTill" db:"active_till"`
	Price              int64     `json:"price"`
	MaxInsurancePayoff int64     `json:"maxInsurancePayoff"`
	MinInsurancePayoff int64     `json:"minInsurancePayoff"`
	AvgInsurancePayoff int64     `json:"avgInsurancePayoff"`
}

type OutboxItem struct {
	ID      string `bson:"id"`
	GroupID int    `bson:"groupID"`
	Status  int    `bson:"status"`
	Val     any    `bson:"val"`
}
