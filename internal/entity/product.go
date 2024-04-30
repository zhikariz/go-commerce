package entity

import "github.com/google/uuid"

type Product struct {
	ID       uuid.UUID
	Name     string
	Price    int64
	Category string
	Auditable
}
