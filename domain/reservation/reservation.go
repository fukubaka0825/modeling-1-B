package reservation

import (
	"../room"
)

// Reservation 和名: 予約
type Reservation struct {
	roomName          *room.Name
	reservationPeriod *Period
}

// NewReservation : 予約はこの関数を通してインスタンス化します
func NewReservation(roomName *room.Name, reservationPeriod *Period) *Reservation {
	return &Reservation{roomName, reservationPeriod}
}

// RoomName getter
func (reservation *Reservation) RoomName() *room.Name {
	return reservation.roomName
}

// ReservationPeriod getter
func (reservation *Reservation) Period() *Period {
	return reservation.reservationPeriod
}
