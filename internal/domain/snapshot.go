package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type Snapshot struct {

	// Currency from patreon.com for timestamp
	Currency Currency

	// Earnings total in specified currency for timestamp
	MonthlyEarnings decimal.Decimal

	// Patrons count for timestamp
	PatronsCount uint64

	CreatedAt time.Time
}
