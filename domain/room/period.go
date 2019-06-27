package room

import (
	"time"
)

// MinutesPerPeriod 最小予約時間(分)
const MinutesPerPeriod = 30

// Period 最小予約単位
type Period struct {
	from     time.Time
	to       time.Time
	priority Priority
	isVacant bool
}

// NewPeriod コンストラクタ
func NewPeriod(from time.Time, priority Priority, isVacant bool) *Period {
	to := from.Add(MinutesPerPeriod * time.Minute)
	return &Period{from, to, priority, isVacant}
}
