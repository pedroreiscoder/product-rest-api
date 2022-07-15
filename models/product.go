package models

type Product struct {
	ID    uint64  `json:"id"`
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}
