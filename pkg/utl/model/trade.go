package gorsk

import "time"

type Trade struct {
	ID              int
	TargetRisk      float64   `json:"target_risk"`
	Atr             ATR       `json:"atr"`
	Balance         float64   `json:"balance"`
	Pair            string    `json:"pair"`
	Exchange        string    `json:"exchange"`
	Side            string    `json:"side"`
	InitialPosition float64   `json:"initial_position"`
	ActualPosition  float64   `json:"actual_position"`
	TotalPosition   float64   `json:"total_position"`
	EntryPrice      float64   `json:"entry_price"`
	EntryTimestamp  time.Time `json:"entry_timestamp"`
	IsOpen          bool      `json:"is_open"`
	IsClose         bool      `json:"is_close"`
	InitialStopLoss float64   `json:"initial_stop_loss"`
	ActualStopLoss  float64   `json:"actual_stop_loss"`
	ExitPrice       float64   `json:"exit_price"`
	ExitTimestamp   time.Time `json:"exit_timestamp"`
	InitialRisk     float64   `json:"initial_risk"`
	ActualRisk      float64   `json:"actual_risk"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	ChangeHistory   []string  `json:"change_history"`
	IsSplit         bool      `json:"is_split"`
}

type ATR struct {
	Value      float64
	Period     int
	Multiplier int
}
