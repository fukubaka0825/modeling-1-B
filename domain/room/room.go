package room

import "time"

// Room 和名: 会議室
type Room struct {
	roomName          *Name
	capacity          int
	reservablePeriods *ReservablePeriods
}

// NewRoom コンストラクタ
func NewRoom(roomName *Name, capacity int, startTime time.Time, finishTime time.Time) *Room {
	// TODO
	return nil
}

// Reserve 予約を確保する
func (room *Room) Reserve( /*Reservationで渡すと循環参照になる。バラで渡す?パッケージわけしないほうがいい？*/ ) bool {
	// TODO
	return true
}
