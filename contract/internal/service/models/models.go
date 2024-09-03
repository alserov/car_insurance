package models

import "time"

const (
	GroupNewInsurances = iota
)

const (
	Succeeded = iota
	Pending
	Failed
)

type (
	OutboxItem struct {
		ID      string `json:"id"`
		GroupID int    `json:"groupID"`
		Status  int    `json:"status"`
		Val     any    `json:"val"`
	}

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
