package models

import "time"

type (
	NewInsurance struct {
		Sender     string    `json:"sender"`
		Amount     int64     `json:"amount"`
		ActiveTill time.Time `json:"activeTill"`
	}
	Payoff struct {
		Receiver string `json:"receiver"`
		Mult     int64  `json:"mult"`
	}
)
