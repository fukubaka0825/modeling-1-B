package room

import "../../domain"

// Room 和名: 会議室
type Room struct {
	roomName         *Name
	reservablePeriod *ReservablePeriod
}

// NewRoom : 会議室はこの関数を通してインスタンス化します
func NewRoom(roomName *Name) *Room {
	return &Room{roomName, NewReservablePeriod()}
}

// CanNotReserve 予約可能時間内ではないか
func (room *Room) CanNotReserve(perod *domain.Period) bool {
	// あとでがんばる
	return true
}
