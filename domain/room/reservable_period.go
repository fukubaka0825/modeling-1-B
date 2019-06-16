package room

import (
	"../../domain"
	_ "time"
)

// ReservablePeriod 和名: 予約可能時間
type ReservablePeriod struct {
}

// NewReservablePeriod : 予約可能時間はこの関数を通してインスタンス化します
func NewReservablePeriod() *ReservablePeriod {
	return &ReservablePeriod{ 
		// &from{time.Date(2018, 3, 3, 17, 44, 13, 0, time.Local)},
		// &to{time.Date(2018, 3, 3, 17, 44, 13, 0, time.Local)},
	}
}

// IsWithIn 予約可能時間内か
func (reservablePeriod *ReservablePeriod) IsWithIn(period domain.Period) bool {
	// あとでがんばる
	return true
}
