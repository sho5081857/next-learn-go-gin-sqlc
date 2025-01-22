// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Customer struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	ImageUrl string    `json:"image_url"`
}

type Invoice struct {
	ID         uuid.UUID   `json:"id"`
	CustomerID uuid.UUID   `json:"customer_id"`
	Amount     int32       `json:"amount"`
	Status     string      `json:"status"`
	Date       pgtype.Date `json:"date"`
}

type Revenue struct {
	Month   string `json:"month"`
	Revenue int32  `json:"revenue"`
}

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
