package reservation

import (
	"time"
	"../../domain"
)

// Period 和名: 予約日時
type Period struct {
	from time.Time
	to   int
}

// ToPeriod domain.Periodに変換します
func (reservationPeriod *Period) ToPeriod() *domain.Period {
	return &domain.Period{
		// あとでがんばる
	}
}
