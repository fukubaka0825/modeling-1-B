package reservation

// ListRepository 和名: 予約リストRepository
type ListRepository interface {
	Record(reservation *Reservation)
}
