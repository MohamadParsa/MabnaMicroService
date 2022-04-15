// Package trade holds entities of trade.
package trade

import "time"

// Trade contain trade data.
type Trade struct {
	TradeID     int32     `json:"trade_id"`
	InstumentID int32     `json:"instument_id"`
	DateEN      time.Time `json:"date_en"`
	Open        float64   `json:"open"`
	High        float64   `json:"high"`
	Low         float64   `json:"low"`
	Close       float64   `json:"close"`
}
