package room

import (
	"time"
)

// 最小予約単位(分)
const MinutesPerPeriod = 60

// Period 和名: 予約日時
type Period struct {
	from     time.Time
	to       time.Time
	priority Priority
	isVacant bool
}

// 予約時間
func NePeriod(from time.Time, priority Priority, isVacant bool) *Period {
	to := from.Add(MinutesPerPeriod * time.Minute)
	return &Period{from, to, priority, isVacant}
}
