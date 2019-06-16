package reservatoinservice

import (
	"../db/reservation"
	"../domain/reservation"
	"../domain/room"
	"errors"
)

// ReservationDetial 予約の詳細
type ReservationDetial struct {
}

var listRepository reservation.ListRepository = db.NewListRepositoryImpl()
// Reserve service
func Reserve(reservation *reservation.Reservation) (*ReservationDetial, *reservation.Reservation, error) {
	// 1. 予約判定をする
	deplecatedReservation, err := canNotReserve(reservation.RoomName(), reservation.Period())
	if err != nil {
		return nil, nil, err
	}
	if deplecatedReservation != nil {
		return nil, deplecatedReservation, nil
	}

	// 2. 予約する
	listRepository.Record(reservation)

	return new(ReservationDetial), nil, nil
}

func canNotReserve(roomName *room.Name, period *reservation.Period) (deplecatedReservation *reservation.Reservation, err error) {
	// 利用可能時間確認
	if room.NewRoom(roomName).CanNotReserve(period.ToPeriod()) {
		return nil, errors.New("利用可能時間確認失敗")
	}
	// 重複確認
	// if {
	//  return false
	// }
	return nil, nil
}
