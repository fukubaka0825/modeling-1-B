package reservation

import (
	"time"
	"fmt"
	"github.com/google/uuid"
	"github.com/takafk9/modeling-1-B/domain/room"
)


// 受け付ける時刻フォーマット（仮置き)
const DatetimeFormat = "2006-01-02 15:04:05"

// 最小予約単位(分)　ここでいい？ -> Periodに持つべきか
const MinutesPerUnit = 60

// Reservation 和名: 予約
type Reservation struct {
	id          uuid.UUID
	roomName    *room.Name
	numOfPeople int
	startTime   *time.Time
	finishTime  *time.Time
	reservedBy	string
	priority room.Priority
}

// NewReservation : 予約はこの関数を通してインスタンス化します
func NewReservation(
	roomName *room.Name,
	startTimeJstString string,
	finishTimeJstString string,
	reservedBy string,
	priority room.Priority
	) *Reservation {
	id, _ := uuid.NewUUID()

	startTime, err := time.Parse(DatetimeFormat, startTimeJstString)
	if err != nil  {
		return nil, fmt.Errorf("Parsing startTimeJstString failed: %v", err)
	}

	finishTime, err := time.Parse(DatetimeFormat, startTimeJstString)
	if err != nil {
		return nil, fmt.Errorf("Parsing finishTimeJstString failed: %v", err)
	}

	// 書いておいてなんだが、多分このバリデーションはここでやるべきではない。
	if startTime.Minute % MinutesPerUnit != 0 && endTime.Minute % MinutesPerUnit != 0 {
		return nil, fmt.Errorf("Reservable time must start every %v minutes from *:00", MinutesPerUnit)
	}

	return &Reservation{
		id,
		roomName,
		startTime,
		finishTime,
		reservedBy,
		priority
	}
}


func (reservation *Reservation) Id() {
	return reservation.Id
}

func (reservation *Reservation) RoomName() {
	return reservation.roomName
}

func (reservation *Reservation) Period() *Period {
	return reservation.reservationPeriod
}

func (reservation *Reservation) StartTime() {
	return reservation.startTime
}

func (reservation *Reservation) FinishTime() {
	return reservation.finishTime
}

func (reservation *Reservation) ReservedBy() {
	return reservation.reservedBy
}

func (reservation *Reservation) Priority() {
	return reservation.priority
}