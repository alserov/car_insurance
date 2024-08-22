package models

import "time"

type Insurance struct {
	SenderAddr string    `json:"senderAddr"`
	Amount     int64     `json:"amount"`
	CarImage   []byte    `json:"carImage"`
	ActiveTill time.Time `json:"activeTill"`
}

type Payoff struct {
	CarImage     []byte  `json:"carImage"`
	ReceiverAddr string  `json:"receiverAddr"`
	Multiplier   float32 `json:"multiplier"`
}

type InsuranceData struct {
	Status             int       `json:"status"`
	ActiveTill         time.Time `json:"activeTill"`
	Owner              string    `json:"owner"`
	Price              int64     `json:"price"`
	MaxInsurancePayoff int64     `json:"maxInsurancePayoff"`
	MinInsurancePayoff int64     `json:"minInsurancePayoff"`
	AvgInsurancePayoff int64     `json:"avgInsurancePayoff"`
}
