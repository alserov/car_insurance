package models

type Insurance struct {
	SenderAddr string `json:"senderAddr"`
	Amount     int64  `json:"amount"`
	CarImage   []byte `json:"carImage"`
}

type Payoff struct {
	CarImage     []byte `json:"carImage"`
	ReceiverAddr string `json:"receiverAddr"`
}

type InsuranceData struct {
	Status             int    `json:"status"`
	ActiveTill         string `json:"activeTill"`
	Owner              string `json:"owner"`
	Price              int64  `json:"price"`
	MaxInsurancePayoff int64  `json:"maxInsurancePayoff"`
	MinInsurancePayoff int64  `json:"minInsurancePayoff"`
	AvgInsurancePayoff int64  `json:"avgInsurancePayoff"`
}
