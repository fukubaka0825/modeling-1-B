package room

import "time"

// 予約枠一覧
type ReservablePeriods struct {
	periodMap map[time.Time]*Period
}

// 予約枠一覧を新規作成
func NewReservablePeriods(from time.Time, to time.Time) *ReservablePeriods {
	// TODO
	return nil
}

// 予約自体の永続化はこのレベル？
