package reservation

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/takafk9/modeling-1-B/domain/room"
)

// DatetimeFormat リクエストパラメータにおけるフォーマット
const DatetimeFormat = "2006-01-02 15:04:05"

// MinutesPerUnit 最小予約単位(分)　ここでいい？ -> Periodに持つべきか
const MinutesPerUnit = 60

// Reservation 予約
type Reservation struct {
	id          uuid.UUID
	roomName    room.Name
	numOfPeople int
	startTime   *time.Time
	finishTime  *time.Time
	reservedBy  string
	priority    room.Priority
}

// NewReservation コンストラクタ
func NewReservation(
	roomName room.Name,
	startTimeJstString string,
	finishTimeJstString string,
	reservedBy string,
	priority room.Priority) (*Reservation, error) {
	id, _ := uuid.NewUUID()

	startTime, err := time.Parse(DatetimeFormat, startTimeJstString)
	if err != nil {
		return nil, fmt.Errorf("Parsing startTimeJstString failed: %v", err)
	}

	finishTime, err := time.Parse(DatetimeFormat, startTimeJstString)
	if err != nil {
		return nil, fmt.Errorf("Parsing finishTimeJstString failed: %v", err)
	}

	// 書いておいてなんだが、多分このバリデーションはここでやるべきではない。
	if startTime.Minute()%MinutesPerUnit != 0 && finishTime.Minute()%MinutesPerUnit != 0 {
		return nil, fmt.Errorf("Reservable time must start every %v minutes from *:00", MinutesPerUnit)
	}

	return &Reservation{
		id:         id,
		roomName:   roomName,
		startTime:  &startTime,
		finishTime: &finishTime,
		reservedBy: reservedBy,
		priority:   priority,
	}, nil
}

// ID 予約ID
func (reservation *Reservation) ID() uuid.UUID {
	return reservation.id
}

// RoomName 部屋名
func (reservation *Reservation) RoomName() room.Name {
	return reservation.roomName
}

// StartTime 開始時間
func (reservation *Reservation) StartTime() *time.Time {
	return reservation.startTime
}

// FinishTime 終了時間
func (reservation *Reservation) FinishTime() *time.Time {
	return reservation.finishTime
}

// ReservedBy 予約者
func (reservation *Reservation) ReservedBy() string {
	return reservation.reservedBy
}

// Priority 優先度
func (reservation *Reservation) Priority() room.Priority {
	return reservation.priority
}
