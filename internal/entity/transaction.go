package entity

import "github.com/google/uuid"

type Transaction struct {
	ID        uuid.UUID
	ProductID uuid.UUID
	Qty       int
	UserID    uuid.UUID
	Discount  *int64
	IsPaid    bool
	Auditable
}
