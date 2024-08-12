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
