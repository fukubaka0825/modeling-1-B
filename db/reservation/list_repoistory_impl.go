package db

import "../../domain/reservation"

// ListRepositoryImpl 実装くらす
type ListRepositoryImpl struct {
}

// Record db層なのでカラッポ!
func (listRepo *ListRepositoryImpl) Record(reservation *reservation.Reservation) {

}

// NewListRepositoryImpl  コンストラクタ
func NewListRepositoryImpl() *ListRepositoryImpl {
	return &ListRepositoryImpl{}
}
