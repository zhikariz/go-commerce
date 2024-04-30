package entity

import "github.com/google/uuid"

type Shipping struct {
	ID            uuid.UUID
	TransactionID uuid.UUID
	IsShipped     bool
	ShippingCost  *int64
	Auditable
}
