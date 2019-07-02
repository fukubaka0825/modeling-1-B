package reservation

import (
	"github.com/google/uuid"
)

// ReservationRepositoryのI/F
type IReservation struct{
	Load(id uuid.UUID) *Reservation

	Save(reservation *Reservation)

}