package models

type (
	NewInsurance struct {
		Sender string `json:"sender"`
		Amount int64  `json:"amount"`
	}
	Payoff struct {
		Receiver string `json:"receiver"`
		Mult     int64  `json:"mult"`
	}
)
